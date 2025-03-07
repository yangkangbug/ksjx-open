package config

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

var nacosConfigClient config_client.IConfigClient

func initNacosConfigClient() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "console.nacos.io",
			ContextPath: "/nacos",
			Port:        80,
		},
	}

	var err error
	nacosConfigClient, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create Nacos config client: %v", err)
	}
}

func loadNacosConfig(dataId, group string) string {
	content, err := nacosConfigClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		log.Printf("Failed to load Nacos config: %v", err)
		return ""
	}
	return content
}
