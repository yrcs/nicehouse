package util

import (
	"strconv"
)

import (
	"github.com/gin-gonic/gin"
)

import (
	"github.com/yrcs/nicehouse/third_party/common"
)

func PackPagingData(ctx *gin.Context, req *common.PagingRequest) {
	pInt, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	psInt, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "1000"))
	page := uint32(pInt)
	pageSize := uint32(psInt)
	query := ctx.QueryMap("query")
	req.Page = &page
	req.PageSize = &pageSize
	req.Query = query
	orderBy := ctx.QueryMap("orderBy")
	for k, v := range orderBy {
		ov, _ := strconv.Atoi(v)
		switch ov {
		case 0:
			req.OrderBy[k] = common.Order_ASC
		case 1:
			req.OrderBy[k] = common.Order_DESC
		default:
			req.OrderBy[k] = common.Order_ASC
		}
	}
}
