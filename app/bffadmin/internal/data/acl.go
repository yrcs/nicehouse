package data

import (
	"context"
)

import (
	aclv1 "github.com/yrcs/nicehouse/api/acl/v1"
	"github.com/yrcs/nicehouse/app/bffadmin/internal/biz"
	"github.com/yrcs/nicehouse/third_party/common"
)

var _ biz.ACLRepo = (*aclRepo)(nil)

type aclRepo struct {
	data *Data
}

func NewACLRepo(data *Data) biz.ACLRepo {
	return &aclRepo{data: data}
}

func (r *aclRepo) ListRoles(ctx context.Context, in *common.PagingRequest) (*common.PagingResponse, error) {
	out, err := r.data.ac.ListRoles(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (r *aclRepo) GetRole(ctx context.Context, o *biz.Role) (*biz.Role, error) {
	out, err := r.data.ac.GetRole(ctx, &aclv1.GetRoleRequest{Id: o.Id})
	if err != nil {
		return nil, err
	}
	return &biz.Role{
		Id:       out.Id,
		IsSystem: out.IsSystem,
	}, err
}

func (r *aclRepo) CreateRole(ctx context.Context, o *biz.Role) (*biz.Role, error) {
	out, err := r.data.ac.CreateRole(ctx, &aclv1.CreateRoleRequest{
		Name:        *o.Name,
		Description: o.Description,
	})
	if err != nil {
		return nil, err
	}
	return &biz.Role{
		Id:        out.Id,
		CreatedAt: out.CreatedAt.AsTime(),
		UpdatedAt: out.GetUpdatedAt().AsTime(),
	}, err
}

func (r *aclRepo) UpdateRole(ctx context.Context, o *biz.Role) (*biz.Role, error) {
	out, err := r.data.ac.UpdateRole(ctx, &aclv1.UpdateRoleRequest{
		Id:          o.Id,
		Name:        o.Name,
		Description: o.Description,
	})
	if err != nil {
		return nil, err
	}
	return &biz.Role{
		Id:        out.Id,
		UpdatedAt: out.GetUpdatedAt().AsTime(),
	}, err
}

func (r *aclRepo) DeleteRoles(ctx context.Context, ids []string) error {
	_, err := r.data.ac.DeleteRoles(ctx, &common.CommonDeletesRequest{Ids: ids})
	return err
}
