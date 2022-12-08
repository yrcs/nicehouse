package data

import (
	"github.com/yrcs/nicehouse/app/acl/internal/biz"
	"github.com/yrcs/nicehouse/pkg/repo"
)

var (
	_  biz.RoleRepo = (*roleRepo)(nil)
	br *repo.BaseRepo[biz.E, biz.T]
)

type roleRepo struct {
	*repo.BaseRepo[biz.E, biz.T]
	data *Data
}

func NewRoleRepo(data *Data) biz.RoleRepo {
	br := &repo.BaseRepo[biz.E, biz.T]{
		DB: data.db,
	}
	return &roleRepo{BaseRepo: br, data: data}
}
