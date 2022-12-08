//go:build wireinject
// +build wireinject

// The "//go:build wireinject" tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
)

import (
	"github.com/yrcs/nicehouse/app/acl/internal/biz"
	"github.com/yrcs/nicehouse/app/acl/internal/data"
	"github.com/yrcs/nicehouse/app/acl/internal/service"
)

// wireApp init dubbogo application.
func wireApp(conf map[string]any) (*service.ACLProvider, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet))
}
