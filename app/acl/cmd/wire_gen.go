// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

import (
	"github.com/yrcs/nicehouse/app/acl/internal/biz"
	"github.com/yrcs/nicehouse/app/acl/internal/data"
	"github.com/yrcs/nicehouse/app/acl/internal/service"
)

// Injectors from wire.go:
// wireApp init dubbogo application.
func wireApp(conf map[string]any) (*service.ACLProvider, func(), error) {
	db := data.NewDB(conf)
	dataData, cleanup, err := data.NewData(db)
	if err != nil {
		return nil, nil, err
	}
	roleRepo := data.NewRoleRepo(dataData)
	roleUsecase := biz.NewRoleUsecase(roleRepo)
	aclProvider := service.NewACLProvider(roleUsecase)
	return aclProvider, func() {
		cleanup()
	}, nil
}
