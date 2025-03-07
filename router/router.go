package router

import (
	"ksjx-open/config"
	"ksjx-open/middleware"
	"ksjx-open/proxy"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RouteConfig 定义动态路由配置项
type RouteConfig struct {
	Method   string
	Path     string
	Upstream string
}

// InitRouter 初始化 Gin 路由，并加载动态路由配置
func InitRouter(conf *config.Config) *gin.Engine {
	// 创建 Gin 实例
	r := gin.New()

	// 全局中间件：指标采集、异常恢复
	r.Use(middleware.MetricsMiddleware())
	r.Use(gin.Recovery())

	// 添加公共健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// 创建 API 路由组，添加 JWT 认证中间件
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.JWTAuth(conf))
	{
		// 示例接口：用户信息（需要根据实际业务扩展）
		apiGroup.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "用户信息"})
		})
	}

	// 加载动态路由（模拟从配置中心获取配置）
	routes := loadDynamicRoutes(conf)
	for _, route := range routes {
		r.Handle(route.Method, route.Path, proxy.CreateHandler(route.Upstream, conf))
	}

	return r
}

// loadDynamicRoutes 模拟从配置中心或 Apollo 获取动态路由配置
func loadDynamicRoutes(conf *config.Config) []RouteConfig {
	return []RouteConfig{
		{Method: "GET", Path: "/service1", Upstream: "service1"},
		{Method: "POST", Path: "/service2", Upstream: "service2"},
	}
}
