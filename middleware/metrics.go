package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// 定义 Prometheus 指标
var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "总 HTTP 请求数",
		},
		[]string{"method", "path", "status"},
	)

	responseTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_time_seconds",
			Help:    "HTTP 响应时间分布",
			Buckets: []float64{0.1, 0.5, 1, 2, 5},
		},
		[]string{"method", "path"},
	)
)

func init() {
	// 注册指标
	prometheus.MustRegister(requestsTotal, responseTime)
}

// MetricsMiddleware 统计请求指标
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())
		requestsTotal.WithLabelValues(c.Request.Method, c.Request.URL.Path, status).Inc()
		responseTime.WithLabelValues(c.Request.Method, c.Request.URL.Path).Observe(duration)
	}
}
