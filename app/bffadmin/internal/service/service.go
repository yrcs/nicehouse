package service

import (
	"github.com/google/wire"
)

import (
	v1 "github.com/yrcs/nicehouse/api/bffadmin/v1"
	"github.com/yrcs/nicehouse/app/bffadmin/internal/biz"
)

var ProviderSet = wire.NewSet(NewBFFAdmin)

type BFFAdmin struct {
	v1.UnimplementedBFFAdminServer

	ac *biz.ACLUsecase
}

func NewBFFAdmin(ac *biz.ACLUsecase) *BFFAdmin {
	return &BFFAdmin{ac: ac}
}
