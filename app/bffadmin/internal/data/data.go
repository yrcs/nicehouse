package data

import (
	"github.com/google/wire"
)

import (
	aclv1 "github.com/yrcs/nicehouse/api/acl/v1"
)

var ProviderSet = wire.NewSet(NewData, NewACLRepo)

type Data struct {
	ac *aclv1.ACLClientImpl
}

func NewData(ac *aclv1.ACLClientImpl) *Data {
	return &Data{ac: ac}
}
