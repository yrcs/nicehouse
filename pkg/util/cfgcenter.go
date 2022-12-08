package util

import (
	"os"
	"strconv"
	"strings"
)

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	"gopkg.in/yaml.v3"
)

const configFileEnvKey = "DUBBO_GO_CONFIG_PATH" // key of environment variable dubbogo configure file path

func loadConfig() (dubboConfig map[string]any, err error) {
	configFilePath := "../conf/dubbogo.yaml"
	if configFilePathFromEnv := os.Getenv(configFileEnvKey); configFilePathFromEnv != "" {
		configFilePath = configFilePathFromEnv
	}

	var configFile *os.File
	configFile, err = os.Open(configFilePath)
	if err != nil {
		return
	}

	if err = yaml.NewDecoder(configFile).Decode(&dubboConfig); err != nil {
		return
	}
	return
}

func getClient() (dataId, group string, client config_client.IConfigClient, err error) {
	dubboConfig, err := loadConfig()
	if err != nil {
		return
	}
	rootConfig := dubboConfig["dubbo"].(map[string]any)
	cfgCenterConfig := rootConfig["config-center"].(map[string]any)
	address := cfgCenterConfig["address"].(string)
	pos := strings.Index(address, ":")
	ipAddr := address[:pos]
	port, _ := strconv.ParseUint(address[pos+1:], 10, 64)
	namespaceId := cfgCenterConfig["namespace"].(string)
	dataId = cfgCenterConfig["data-id"].(string)
	group = cfgCenterConfig["group"].(string)

	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(ipAddr, port, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(namespaceId),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	// create config client
	client, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	return
}

func GetCustomConfig() (map[string]any, error) {
	dataId, group, client, err := getClient()
	if err != nil {
		return nil, err
	}

	// get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return nil, err
	}

	var m map[string]any
	err = yaml.Unmarshal([]byte(content), &m)
	if err != nil {
		return nil, err
	}

	rootConfig := m["dubbo"].(map[string]any)
	customConfig := rootConfig["custom"].(map[string]any)
	return customConfig["config-map"].(map[string]any), err
}
