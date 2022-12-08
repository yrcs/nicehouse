//go:build wireinject
// +build wireinject

// The "//go:build wireinject" tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
)

import (
	aclv1 "github.com/yrcs/nicehouse/api/acl/v1"
	"github.com/yrcs/nicehouse/app/bffadmin/internal/biz"
	"github.com/yrcs/nicehouse/app/bffadmin/internal/data"
	"github.com/yrcs/nicehouse/app/bffadmin/internal/service"
)

// wireApp init dubbogo application.
func wireApp(ac *aclv1.ACLClientImpl) *service.BFFAdmin {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet))
}
