package service

import (
	"context"
	"errors"
)

import (
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gorm.io/gorm"
)

import (
	v1 "github.com/yrcs/nicehouse/api/acl/v1"
	"github.com/yrcs/nicehouse/app/acl/internal/biz"
	"github.com/yrcs/nicehouse/app/acl/internal/biz/do"
	"github.com/yrcs/nicehouse/pkg/pagination"
	"github.com/yrcs/nicehouse/pkg/result"
	"github.com/yrcs/nicehouse/pkg/usecase"
	"github.com/yrcs/nicehouse/pkg/util"
	"github.com/yrcs/nicehouse/third_party/common"
)

func (s *ACLProvider) ListRoles(ctx context.Context, in *common.PagingRequest) (*common.PagingResponse, error) {
	offset, limit, orderBy := pagination.GetPagingParams(in.GetPage(), in.GetPageSize(), in.OrderBy)
	var (
		roles []biz.E
		total int
		err   error
	)
	roles, total, err = s.rc.ListByPage(ctx, offset, limit, orderBy, "Name LIKE ?", "%"+in.Query["Name"]+"%")
	if err != nil {
		return nil, result.ErrorWithDetails(ctx, nil, err)
	}
	items := make([]*anypb.Any, 0, len(roles))
	for _, r := range roles {
		role := &v1.Role{
			Id:          r.Id,
			CreatedAt:   timestamppb.New(r.CreatedAt),
			UpdatedAt:   timestamppb.New(r.UpdatedAt),
			Name:        r.Name,
			Description: r.Description,
			IsSystem:    r.IsSystem,
		}
		a, _ := anypb.New(role)
		items = append(items, a)
	}
	return &common.PagingResponse{
		Total: uint32(total),
		Items: items,
	}, nil
}

func (s *ACLProvider) GetRole(ctx context.Context, in *v1.GetRoleRequest) (*v1.Role, error) {
	out, err := s.rc.Get(ctx, "Id = ?", in.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, result.ErrorWithDetails(ctx, biz.ErrRoleNotFound, err)
		}
		result.ErrorWithDetails(ctx, nil, err)
	}
	if out.IsSystem {
		return nil, result.ErrorWithDetails(ctx, biz.ErrRoleUnprocessableEntity, err)
	}
	return &v1.Role{
		Id:       out.Id,
		IsSystem: out.IsSystem,
	}, nil
}

func (s *ACLProvider) CreateRole(ctx context.Context, in *v1.CreateRoleRequest) (*common.CommonCreate, error) {
	id, err := util.MakeULID()
	if err != nil {
		return nil, result.ErrorWithDetails(ctx, nil, err)
	}
	out, err := s.rc.Create(ctx, &do.Role{
		BaseDO:      usecase.BaseDO{Id: id},
		Name:        in.Name,
		Description: in.GetDescription(),
	})
	if err != nil {
		return nil, result.ErrorWithDetails(ctx, nil, err)
	}
	return &common.CommonCreate{
		Id:        out.Id,
		CreatedAt: timestamppb.New(out.CreatedAt),
		UpdatedAt: timestamppb.New(out.UpdatedAt),
	}, nil
}

func (s *ACLProvider) UpdateRole(ctx context.Context, in *v1.UpdateRoleRequest) (*common.CommonUpdate, error) {
	values := make(map[string]any, 3)
	util.UpdateOptionalFields(in, values)
	out, err := s.rc.Updates(ctx, values)
	if err != nil {
		return nil, result.ErrorWithDetails(ctx, nil, err)
	}
	return &common.CommonUpdate{
		Id:        out.Id,
		UpdatedAt: timestamppb.New(out.UpdatedAt),
	}, nil
}

func (s *ACLProvider) DeleteRoles(ctx context.Context, in *common.CommonDeletesRequest) (*emptypb.Empty, error) {
	if len(in.Ids) < 1 {
		return nil, nil
	}
	if err := s.rc.Delete(ctx, in.Ids, "IsSystem = ?", false); err != nil {
		return nil, result.ErrorWithDetails(ctx, nil, err)
	}
	return nil, nil
}
