package utils

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

var nacosClient naming_client.INamingClient

func InitNacosClient() {
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
	nacosClient, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create Nacos client: %v", err)
	}
}

func GetServiceInstances(serviceName string) []string {
	instances, err := nacosClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: serviceName,
	})
	if err != nil {
		log.Printf("Failed to get service instances from Nacos: %v", err)
		return nil
	}

	var instanceAddresses []string
	for _, instance := include instances {
		instanceAddresses = append(instanceAddresses, instance.Ip)
	}
	return instanceAddresses
}
