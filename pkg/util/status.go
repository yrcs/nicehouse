package util

import (
	"context"
	"fmt"
	"net/http"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/constant"

	"github.com/dubbogo/grpc-go/codes"
	"github.com/dubbogo/grpc-go/status"

	httpstatus "github.com/go-kratos/kratos/v2/transport/http/status"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func NewStatus(c int, msg string) *status.Status {
	switch c {
	case http.StatusUnprocessableEntity:
		return status.New(codes.Code(c), msg)
	}
	return status.New(codes.Code(httpstatus.ToGRPCCode(c)), msg)
}

func StatusError(ctx context.Context, err error) error {
	m := make(map[string]string, 1)
	atm := ctx.Value(constant.AttachmentKey).(map[string]any)
	m[constant.InterfaceKey] = atm[constant.InterfaceKey].([]string)[0]
	statusWithDetails, err := status.Convert(err).WithDetails(
		&errdetails.ErrorInfo{
			Metadata: m,
		},
		&errdetails.BadRequest_FieldViolation{
			Field:       "cause",
			Description: fmt.Sprintf("%+v", err),
		},
	)
	if err != nil {
		return status.Errorf(codes.Unknown, "WithDetails error: %+v interface: %s", err, m[constant.InterfaceKey])
	}
	return statusWithDetails.Err()
}

func StatusWithDetails(ctx context.Context, s *status.Status, reason string, e error, metadata ...map[string]string) error {
	m := make(map[string]string, 1)
	atm := ctx.Value(constant.AttachmentKey).(map[string]any)
	m[constant.InterfaceKey] = atm[constant.InterfaceKey].([]string)[0]
	if len(metadata) > 0 {
		for k, v := range metadata[0] {
			m[k] = v
		}
	}
	statusWithDetails, err := s.WithDetails(
		&errdetails.ErrorInfo{
			Reason:   reason,
			Metadata: m,
		},
		&errdetails.BadRequest_FieldViolation{
			Field:       "cause",
			Description: fmt.Sprintf("%+v", e),
		},
	)
	if err != nil {
		return status.Errorf(codes.Unknown, "WithDetails error: %+v interface: %s", err, m[constant.InterfaceKey])
	}
	return statusWithDetails.Err()
}
