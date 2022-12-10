package v1

import (
	"context"
)

import (
	"github.com/gin-gonic/gin"

	"google.golang.org/protobuf/types/known/emptypb"
)

import (
	"github.com/yrcs/nicehouse/pkg/result"
	"github.com/yrcs/nicehouse/pkg/util"
	"github.com/yrcs/nicehouse/third_party/common"
)

type BFFAdminHTTPServer interface {
	ListRoles(context.Context, *common.PagingRequest) (*common.PagingResponse, error)
	GetRole(context.Context, *GetRoleRequest) (*Role, error)
	CreateRole(context.Context, *CreateRoleRequest) (*common.CommonCreate, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*common.CommonUpdate, error)
	DeleteRoles(context.Context, *common.CommonDeletesRequest) (*emptypb.Empty, error)
}

func RegisterBFFAdminHTTPServer(r *gin.Engine, srv BFFAdminHTTPServer) {
	v1 := r.Group("v1/admin")
	{
		g1 := v1.Group("roles")
		{
			// ?page=1&pageSize=10&query[name]=超级管理员&query[IsSystem]=true&orderBy[name]=1&orderBy[id]=0
			g1.GET("", ListRolesHandler(srv))
			g1.GET(":id", GetRoleHandler(srv))
			g1.POST("", CreateRoleHandler(srv))
			g1.PUT(":id", UpdateRoleHandler(srv))
			g1.DELETE("", DeleteRolesHandler(srv))
		}
	}
}

// @Tags        获取角色管理列表
// @Summary     角色列表
// @Description 角色管理分页列表
// @Accept      json
// @Produce     json
// @Param       page               query    int    false "页码"    Format(uint32)
// @Param       pageSize           query    int    false "每页条目数" Format(uint32)
// @Param       query[Name]        query    string false "名称"
// @Param       query[Description] query    string false "描述"
// @Param       query[IsSystem]    query    bool   false "是否内置"
// @Param       orderBy[Name]      query    int    false "按名称排序"   Enums(0, 1)
// @Param       orderBy[Id]        query    int    false "按 ID 排序" Enums(0, 1)
// @Success     200                {object} common.PagingResponse
// @Router      /admin/roles [get]
func ListRolesHandler(srv BFFAdminHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		o := make(map[string]common.Order, len(ctx.QueryMap("orderBy")))
		in := common.PagingRequest{OrderBy: o}
		util.PackPagingData(ctx, &in)
		out, err := srv.ListRoles(ctx, &in)
		if err != nil {
			result.Result(ctx, err)
			return
		}
		result.Result(ctx, out)
	}
}

// @Tags        获取一个角色
// @Summary     获取一个角色
// @Description 通过接收 id 参数来获取一个角色
// @Accept      json
// @Produce     json
// @Param       id  path     string true "角色 id" minlength(26) maxlength(26)
// @Success     200 {object} v1.Role
// @Router      /admin/roles/{id} [get]
func GetRoleHandler(srv BFFAdminHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		in := GetRoleRequest{Id: id}
		out, err := srv.GetRole(ctx, &in)
		if err != nil {
			result.Result(ctx, err)
			return
		}
		result.Result(ctx, out)
	}
}

// @Tags        添加角色
// @Summary     添加角色
// @Description 通过接收 json body 参数来添加角色
// @Accept      json
// @Produce     json
// @Param       message body     v1.CreateRoleRequest false "角色属性"
// @Success     200     {object} common.CommonCreate
// @Router      /admin/roles [post]
func CreateRoleHandler(srv BFFAdminHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in CreateRoleRequest
		if err := ctx.ShouldBindJSON(&in); err != nil {
			result.Result(ctx, err)
			return
		}
		out, err := srv.CreateRole(ctx, &in)
		if err != nil {
			result.Result(ctx, err)
			return
		}
		result.Result(ctx, out)
	}
}

// @Tags        编辑一个角色
// @Summary     编辑一个角色
// @Description 通过接收 id 参数和 json body 参数来编辑一个角色
// @Accept      json
// @Produce     json
// @Param       id      path     string               true  "角色 id" minlength(26) maxlength(26)
// @Param       message body     v1.UpdateRoleRequest false "角色属性"
// @Success     200     {object} common.CommonUpdate
// @Router      /admin/roles/{id} [put]
func UpdateRoleHandler(srv BFFAdminHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		in := UpdateRoleRequest{Id: id}
		if err := ctx.ShouldBindJSON(&in); err != nil {
			result.Result(ctx, err)
			return
		}
		_, err := srv.GetRole(ctx, &GetRoleRequest{Id: id})
		if err != nil {
			result.Result(ctx, err)
			return
		}
		out, err := srv.UpdateRole(ctx, &in)
		if err != nil {
			result.Result(ctx, err)
			return
		}
		result.Result(ctx, out)
	}
}

// @Tags        删除角色
// @Summary     删除一个或多个角色
// @Description 通过接收数组 ids json body 参数来删除角色
// @Accept      json
// @Produce     json
// @Param       message body common.CommonDeletesRequest true "角色 ids 数组"
// @Success     200
// @Router      /admin/roles [delete]
func DeleteRolesHandler(srv BFFAdminHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in common.CommonDeletesRequest
		if err := ctx.ShouldBindJSON(&in); err != nil {
			result.Result(ctx, err)
			return
		}
		out, err := srv.DeleteRoles(ctx, &in)
		if err != nil {
			result.Result(ctx, err)
			return
		}
		result.Result(ctx, out)
	}
}
