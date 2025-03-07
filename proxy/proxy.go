package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"ksjx-open/config"
	"ksjx-open/service"

	"github.com/gin-gonic/gin"
)

// InitProxy 初始化全局反向代理配置，如自定义 HTTP Transport（此处简化处理）
func InitProxy(conf *config.Config) {
	// 可根据需要自定义 http.DefaultTransport 配置，实现连接池复用
}

// CreateHandler 根据上游服务名称创建反向代理处理函数
func CreateHandler(upstream string, conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		instances := service.GetService(upstream)
		target := service.Select(instances)
		if target == "" {
			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "未找到可用服务"})
			return
		}
		targetURL, err := url.Parse(target)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "服务地址错误"})
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}