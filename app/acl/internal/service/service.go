package service

import (
	"github.com/google/wire"
)

import (
	v1 "github.com/yrcs/nicehouse/api/acl/v1"
	"github.com/yrcs/nicehouse/app/acl/internal/biz"
)

var ProviderSet = wire.NewSet(NewACLProvider)

type ACLProvider struct {
	v1.UnimplementedACLServer

	rc *biz.RoleUsecase
}

func NewACLProvider(rc *biz.RoleUsecase) *ACLProvider {
	return &ACLProvider{rc: rc}
}
