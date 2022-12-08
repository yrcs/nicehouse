package result

import (
	"encoding/json"
	"net/http"
)

import (
	"github.com/dubbogo/gost/log/logger"

	dubbostatus "github.com/dubbogo/grpc-go/status"

	"github.com/gin-gonic/gin"

	"github.com/go-errors/errors"

	kratoserrors "github.com/go-kratos/kratos/v2/errors"
	httpstatus "github.com/go-kratos/kratos/v2/transport/http/status"

	"google.golang.org/genproto/googleapis/rpc/errdetails"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/encoding/protojson"

	"google.golang.org/protobuf/proto"
)

const (
	maxGRPCCode     = 16
	jsonContentType = "application/json; charset=utf-8"
)

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
	se := fromError(err)
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

// fromError try to convert an error to *Error.
// It supports wrapped errors.
func fromError(err error) *kratoserrors.Error {
	if err == nil {
		return nil
	}
	if se := new(kratoserrors.Error); kratoserrors.As(err, &se) {
		return se
	}
	var gs *status.Status
	ds, ok := dubbostatus.FromError(err)
	if !ok {
		if gs, ok = status.FromError(err); !ok {
			return kratoserrors.New(kratoserrors.UnknownCode, kratoserrors.UnknownReason, err.Error())
		}
	}
	var (
		code    codes.Code
		message string
		details []any
	)
	if ds != nil {
		code = codes.Code(ds.Code())
		message = ds.Message()
		details = ds.Details()
	} else {
		code = gs.Code()
		message = gs.Message()
		details = gs.Details()
	}
	var httpCode int
	if code > maxGRPCCode {
		httpCode = int(code)
	} else {
		httpCode = httpstatus.FromGRPCCode(code)
	}
	ret := kratoserrors.New(
		httpCode,
		kratoserrors.UnknownReason,
		message,
	)
	for _, detail := range details {
		switch d := detail.(type) {
		case *errdetails.ErrorInfo:
			ret.Reason = d.Reason
			ret = ret.WithMetadata(d.Metadata)
		case *errdetails.BadRequest_FieldViolation:
			if d.Field == "cause" {
				ret = ret.WithCause(errors.New(d.Description))
			}
		}
	}
	return ret
}

func marshalJSON(v any) ([]byte, error) {
	if m, ok := v.(proto.Message); ok {
		return protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(m)
	}
	return json.Marshal(v)
}
