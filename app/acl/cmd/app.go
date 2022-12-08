package main

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"github.com/dubbogo/gost/log/logger"

	"github.com/knadh/koanf"
)

import (
	"github.com/yrcs/nicehouse/app/acl/internal/service"
)

var cleanup func()

// export DUBBO_GO_CONFIG_PATH=../conf/direct.yaml
func init() {
	//local direct
	rootConfig := config.NewRootConfigBuilder().Build()
	loaderConf := config.NewLoaderConf()
	koan := config.GetConfigResolver(loaderConf)
	koan = loaderConf.MergeConfig(koan)

	var err error
	if err = koan.UnmarshalWithConf(rootConfig.Prefix(),
		rootConfig, koanf.UnmarshalConf{Tag: "yaml"}); err != nil {
		panic(err)
	}
	if err = rootConfig.Custom.Init(); err != nil {
		panic(err)
	}

	conf := rootConfig.Custom.ConfigMap
	if _, exists := conf["database"].(map[string]any); !exists {
		logger.Fatal("please add the database config first.")
	}

	//config center
	//conf, err := util.GetCustomConfig()
	//if err != nil {
	//	panic(err)
	//}

	var srv *service.ACLProvider
	srv, cleanup, err = wireApp(conf)
	if err != nil {
		panic(err)
	}

	config.SetProviderService(srv)
}

func main() {
	defer cleanup()

	if err := config.Load(); err != nil {
		panic(err)
	}

	select {}
}
