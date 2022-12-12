package biz

import (
	"net/http"
)

import (
	v1 "github.com/yrcs/nicehouse/api/acl/v1"
	"github.com/yrcs/nicehouse/app/acl/internal/biz/do"
	"github.com/yrcs/nicehouse/app/acl/internal/data/po"
	myerrors "github.com/yrcs/nicehouse/pkg/errors"
	"github.com/yrcs/nicehouse/pkg/repo"
	"github.com/yrcs/nicehouse/pkg/usecase"
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
	usecase.BaseUsecase[E, T]
	repo RoleRepo
}

func NewRoleUsecase(repo RoleRepo) *RoleUsecase {
	return &RoleUsecase{
		BaseUsecase: usecase.BaseUsecase[E, T]{repo},
		repo:        repo,
	}
}
