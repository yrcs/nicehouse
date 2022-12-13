package service

import (
	"context"
)

import (
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

import (
	v1 "github.com/yrcs/nicehouse/api/bffadmin/v1"
	"github.com/yrcs/nicehouse/app/bffadmin/internal/biz"
	"github.com/yrcs/nicehouse/pkg/usecase"
	"github.com/yrcs/nicehouse/third_party/common"
)

func (s *BFFAdmin) ListRoles(ctx context.Context, in *common.PagingRequest) (*common.PagingResponse, error) {
	out, err := s.ac.ListRoles(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *BFFAdmin) GetRole(ctx context.Context, in *v1.GetRoleRequest) (*v1.Role, error) {
	out, err := s.ac.GetRole(ctx, &biz.Role{BaseDO: usecase.BaseDO{Id: in.Id}})
	if err != nil {
		return nil, err
	}
	return &v1.Role{
		Id:       out.Id,
		IsSystem: out.IsSystem,
	}, nil
}

func (s *BFFAdmin) CreateRole(ctx context.Context, in *v1.CreateRoleRequest) (*common.CommonCreate, error) {
	out, err := s.ac.CreateRole(ctx, &biz.Role{
		Name:        &in.Name,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}
	return &common.CommonCreate{
		Id:        out.Id,
		CreatedAt: timestamppb.New(out.CreatedAt),
		UpdatedAt: timestamppb.New(out.UpdatedAt),
	}, nil
}

func (s *BFFAdmin) UpdateRole(ctx context.Context, in *v1.UpdateRoleRequest) (*common.CommonUpdate, error) {
	out, err := s.ac.UpdateRole(ctx, &biz.Role{
		BaseDO:      usecase.BaseDO{Id: in.Id},
		Name:        in.Name,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}
	return &common.CommonUpdate{
		Id:        out.Id,
		UpdatedAt: timestamppb.New(out.UpdatedAt),
	}, nil
}

func (s *BFFAdmin) DeleteRoles(ctx context.Context, in *common.CommonDeletesRequest) (*emptypb.Empty, error) {
	return nil, s.ac.DeleteRoles(ctx, in.GetIds())
}
