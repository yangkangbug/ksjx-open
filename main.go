package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"net/http"
)

var nacosClient naming_client.INamingClient

func initNacosClient() {
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

func getServiceInstances(serviceName string) []string {
	instances, err := nacosClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: serviceName,
	})
	if err != nil {
		log.Printf("Failed to get service instances from Nacos: %v", err)
		return nil
	}

	var instanceAddresses []string
	for _, instance := range instances {
		instanceAddresses = append(instanceAddresses, instance.Ip)
	}
	return instanceAddresses
}

func createHandler(upstream string) gin.HandlerFunc {
	return func(c *gin.Context) {
		instances := getServiceInstances(upstream)
		if len(instances) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No available instances"})
			return
		}
		target := instances[0] // Simple load balancing strategy
		c.JSON(http.StatusOK, gin.H{"target": target})
	}
}

func main() {
	initNacosClient()

	r := gin.Default()
	r.GET("/service/:name", func(c *gin.Context) {
		serviceName := c.Param("name")
		handler := createHandler(serviceName)
		handler(c)
	})

	r.Run(":8080")
}
