package main

import (
	"fmt"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"github.com/dubbogo/gost/log/logger"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

import (
	aclv1 "github.com/yrcs/nicehouse/api/acl/v1"
	v1 "github.com/yrcs/nicehouse/api/bffadmin/v1"
	"github.com/yrcs/nicehouse/api/bffadmin/v1/docs"
	_ "github.com/yrcs/nicehouse/pkg/filter/validator"
)

var aclClientImpl = &aclv1.ACLClientImpl{}

// export DUBBO_GO_CONFIG_PATH=../conf/direct.yaml
func init() {
	config.SetConsumerService(aclClientImpl)
}

// @title       Nicehouse (好房) 后台管理系统 API
// @version     1.0
// @description 本文档描述了后台管理系统微服务接口定义。

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	serverConfig, ok := config.GetRootConfig().Custom.ConfigMap["server"].(map[string]any)
	if !ok {
		logger.Fatal("please add the server config first.")
	}

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	srv := wireApp(aclClientImpl)
	v1.RegisterBFFAdminHTTPServer(r, srv)
	err := r.Run(fmt.Sprintf("%s:%d", serverConfig["host"], serverConfig["port"]))
	if err != nil {
		panic(err)
	}

	select {}
}
