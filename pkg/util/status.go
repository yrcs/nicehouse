package util

import (
	"fmt"
	"net/http"
)

import (
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

func StatusWithDetails(s *status.Status, reason string, e error, metadata ...map[string]string) error {
	var m map[string]string
	if len(metadata) > 0 {
		m = metadata[0]
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
		return err
	}
	return statusWithDetails.Err()
}
