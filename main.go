package main

import (
	"ksjx-open/client"
	"log"
	"net/http"
	"time"

	"ksjx-open/config"
	"ksjx-open/middleware"
	"ksjx-open/proxy"
	"ksjx-open/router"
	"ksjx-open/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置信息
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化 Redis 连接池
	service.InitRedisPool(conf.RedisAddr)

	// 初始化反向代理（如果使用反向代理功能）
	proxy.InitProxy(conf)

	// 初始化库存查询 gRPC 客户端，连接 Spring Boot 项目的库存查询服务
	if err := client.InitInventoryClient(conf.RedisAddr); err != nil {
		log.Fatalf("初始化库存查询 gRPC 客户端失败: %v", err)
	}

	// 初始化 Gin 路由
	r := gin.New()
	r.Use(middleware.MetricsMiddleware())
	r.Use(gin.Recovery())

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// 注册已有 API 路由（例如用户管理、动态路由等）
	// 此处省略其他路由注册...

	// 注册查询库存的 HTTP 路由
	router.RegisterInventoryRoute(r)

	// 启动 HTTP 服务
	srv := &http.Server{
		Addr:         conf.ServerAddress,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Printf("服务启动，监听地址：%s", conf.ServerAddress)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
