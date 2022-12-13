package biz

import (
	"context"
)

import (
	"github.com/yrcs/nicehouse/pkg/usecase"
	"github.com/yrcs/nicehouse/third_party/common"
)

type Role struct {
	usecase.BaseDO
	Name        *string
	Description *string
	IsSystem    bool
}

type ACLRepo interface {
	ListRoles(ctx context.Context, in *common.PagingRequest) (*common.PagingResponse, error)
	GetRole(ctx context.Context, o *Role) (*Role, error)
	CreateRole(ctx context.Context, o *Role) (*Role, error)
	UpdateRole(ctx context.Context, o *Role) (*Role, error)
	DeleteRoles(ctx context.Context, ids []string) error
}

type ACLUsecase struct {
	repo ACLRepo
}

func NewACLUsecase(repo ACLRepo) *ACLUsecase {
	return &ACLUsecase{repo: repo}
}

func (ac *ACLUsecase) ListRoles(ctx context.Context, in *common.PagingRequest) (*common.PagingResponse, error) {
	return ac.repo.ListRoles(ctx, in)
}

func (ac *ACLUsecase) GetRole(ctx context.Context, o *Role) (*Role, error) {
	return ac.repo.GetRole(ctx, o)
}

func (ac *ACLUsecase) CreateRole(ctx context.Context, o *Role) (*Role, error) {
	return ac.repo.CreateRole(ctx, o)
}

func (ac *ACLUsecase) UpdateRole(ctx context.Context, o *Role) (*Role, error) {
	return ac.repo.UpdateRole(ctx, o)
}

func (ac *ACLUsecase) DeleteRoles(ctx context.Context, ids []string) error {
	return ac.repo.DeleteRoles(ctx, ids)
}
