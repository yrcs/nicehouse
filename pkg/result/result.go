package result

import (
	"context"
	"encoding/json"
	"net/http"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/constant"

	"github.com/dubbogo/gost/log/logger"

	"github.com/gin-gonic/gin"

	"google.golang.org/protobuf/encoding/protojson"

	"google.golang.org/protobuf/proto"
)

import (
	myerrors "github.com/yrcs/nicehouse/pkg/errors"
)

const (
	maxGRPCCode     = 16
	jsonContentType = "application/json; charset=utf-8"
)

func ErrorWithDetails(ctx context.Context, err *myerrors.Error, cause error, metadata ...map[string]string) error {
	m := make(map[string]string, 1)
	atm := ctx.Value(constant.AttachmentKey).(map[string]any)
	m[constant.InterfaceKey] = atm[constant.InterfaceKey].([]string)[0]
	if len(metadata) > 0 {
		for k, v := range metadata[0] {
			m[k] = v
		}
	}
	if err == nil {
		err = myerrors.Newf(myerrors.UnknownCode, myerrors.UnknownReason, "%+v", cause)
	}
	return err.WithMetadata(m).WithCause(cause)
}

func Result(ctx *gin.Context, v any) {
	if err, ok := v.(error); ok {
		errorEncoder(ctx, err)
		return
	}
	responseEncoder(ctx, v)
}

// responseEncoder encodes the object to the HTTP response.
func responseEncoder(ctx *gin.Context, v any) error {
	if v == nil {
		return nil
	}
	body, err := marshalJSON(v)
	if err != nil {
		return err
	}
	ctx.Writer.Header().Set("Content-Type", jsonContentType)
	_, err = ctx.Writer.Write(body)
	if err != nil {
		return err
	}
	return nil
}

// errorEncoder encodes the error to the HTTP response.
func errorEncoder(ctx *gin.Context, err error) {
	se := myerrors.FromError(err)
	logger.Error(se.Error())
	body, err := marshalJSON(se)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Writer.Header().Set("Content-Type", jsonContentType)
	ctx.Writer.WriteHeader(int(se.Code))
	_, _ = ctx.Writer.Write(body)
}

func marshalJSON(v any) ([]byte, error) {
	if m, ok := v.(proto.Message); ok {
		return protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(m)
	}
	return json.Marshal(v)
}
