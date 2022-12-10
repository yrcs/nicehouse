package validator

import (
	"context"
	"fmt"
	"sync"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"

	"github.com/mitchellh/mapstructure"
)

import (
	myerrors "github.com/yrcs/nicehouse/pkg/errors"
)

const (
	validatorFilterKey = "validator"
	reason             = "VALIDATOR_INVALID_ARGUMENT"
)

var (
	validatorOnce     sync.Once
	validatorInstance *validatorFilter
)

func init() {
	extension.SetFilter(validatorFilterKey, newValidatorFilter)
}

type validatorFilter struct{}

func newValidatorFilter() filter.Filter {
	if validatorInstance == nil {
		validatorOnce.Do(func() {
			validatorInstance = &validatorFilter{}
		})
	}
	return validatorInstance
}

type validator interface {
	Validate() error
}

func (f *validatorFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	var req any
	for _, m := range invocation.Arguments() {
		_ = mapstructure.Decode(m, &req)
	}
	if v, ok := req.(validator); ok {
		if err := v.Validate(); err != nil {
			return &protocol.RPCResult{
				Err: myerrors.BadRequest(reason, err.Error()).
					WithMetadata(map[string]string{
						"invoker": fmt.Sprintf("%v", invoker),
						"method":  invocation.MethodName()}).
					WithCause(err),
			}
		}
	}
	return invoker.Invoke(ctx, invocation)
}

func (f *validatorFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, protocol protocol.Invocation) protocol.Result {
	return result
}
