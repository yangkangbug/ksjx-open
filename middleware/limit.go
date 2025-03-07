package middleware

import (
	"net/http"

	"ksjx-open/config"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// 定义全局令牌桶，限制请求速率（单位：qps 与 burst 可根据需要调整）
var limiter = rate.NewLimiter(1000, 200)

// RateLimit 令牌桶限流中间件
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "请求过多"})
			return
		}
		c.Next()
	}
}

// HystrixWrap 熔断保护中间件包装
func HystrixWrap(serviceName string, conf *config.Config) gin.HandlerFunc {
	hystrix.ConfigureCommand(serviceName, hystrix.CommandConfig{
		Timeout:               conf.HystrixTimeout,
		MaxConcurrentRequests: conf.HystrixMaxConcurrentRequests,
		ErrorPercentThreshold: conf.HystrixErrorPercentThreshold,
	})
	return func(c *gin.Context) {
		err := hystrix.Do(serviceName, func() error {
			c.Next()
			return nil
		}, func(err error) error {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "服务暂不可用"})
			return nil
		})
		if err != nil {
			return
		}
	}
}
