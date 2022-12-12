package service

import (
	"context"
	"errors"
	"strconv"
)

import (
	"github.com/jinzhu/copier"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gorm.io/gorm"
)

import (
	v1 "github.com/yrcs/nicehouse/api/acl/v1"
	"github.com/yrcs/nicehouse/app/acl/internal/biz"
	"github.com/yrcs/nicehouse/app/acl/internal/biz/do"
	"github.com/yrcs/nicehouse/pkg/result"
	"github.com/yrcs/nicehouse/pkg/util"
	"github.com/yrcs/nicehouse/third_party/common"
)

const maxPageSize = 1000

func (s *ACLProvider) ListRoles(ctx context.Context, in *common.PagingRequest) (*common.PagingResponse, error) {
	limit := int(in.GetPageSize())
	if limit == 0 || limit > maxPageSize {
		limit = maxPageSize
	}

	offset := int(in.GetPage())
	if offset > 0 {
		offset = (offset - 1) * limit
	}

	orderBy := make(map[string]string, len(in.OrderBy))
	for k, v := range in.OrderBy {
		orderBy[k] = common.Order_name[int32(v.Number())]
	}

	var (
		queryCopy map[string]any
		roles     []biz.E
		total     int
		err       error
	)
	copier.Copy(&queryCopy, in.GetQuery())
	isSystemKey := "IsSystem"
	if v, exists := queryCopy[isSystemKey]; exists {
		isSystem, err := strconv.ParseBool(v.(string))
		if err != nil {
			return nil, result.ErrorWithDetails(ctx, biz.ErrRoleInvalidArgument, err)
		}
		queryCopy[isSystemKey] = isSystem
	}
	roles, total, err = s.rc.ListByPage(ctx, offset, limit, queryCopy, orderBy)
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
		Id:          id,
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
