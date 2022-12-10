package biz

import (
	"context"
	"net/http"
)

import (
	v1 "github.com/yrcs/nicehouse/api/acl/v1"
	"github.com/yrcs/nicehouse/app/acl/internal/biz/do"
	"github.com/yrcs/nicehouse/app/acl/internal/data/po"
	myerrors "github.com/yrcs/nicehouse/pkg/errors"
	"github.com/yrcs/nicehouse/pkg/repo"
)

var (
	ErrRoleInvalidArgument     = myerrors.BadRequest(v1.ErrorReason_ACL_ROLE_INVALID_ARGUMENT.String(), "请求参数错误")
	ErrRoleNotFound            = myerrors.NotFound(v1.ErrorReason_ACL_ROLE_NOT_FOUND.String(), "角色未找到")
	ErrRoleUnprocessableEntity = myerrors.New(http.StatusUnprocessableEntity, v1.ErrorReason_ACL_ROLE_UNPROCESSABLE_ENTITY.String(), "请求无法处理")
)

type E *do.Role
type T *po.Role

type RoleRepo interface {
	repo.Repo[E, T]
}

type RoleUsecase struct {
	repo RoleRepo
}

func NewRoleUsecase(repo RoleRepo) *RoleUsecase {
	return &RoleUsecase{repo: repo}
}

func (c *RoleUsecase) List(ctx context.Context, offset int, limit int, conds map[string]any, orderBy map[string]string) ([]E, int, error) {
	return c.repo.FindByPage(ctx, offset, limit, conds, orderBy)
}

func (c *RoleUsecase) Get(ctx context.Context, conds ...any) (E, error) {
	return c.repo.FindOne(ctx, conds)
}

func (c *RoleUsecase) Create(ctx context.Context, o E) (E, error) {
	if err := c.repo.Create(ctx, o); err != nil {
		return nil, err
	}
	return o, nil
}

func (c *RoleUsecase) Update(ctx context.Context, values map[string]any) (E, error) {
	id := values["Id"]
	delete(values, "Id")
	o, err := c.repo.Updates(ctx, values, "id = ?", id)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (c *RoleUsecase) Delete(ctx context.Context, ids []string, query any, conds ...any) error {
	return c.repo.Delete(ctx, ids, query, conds...)
}
