package config

import (
	"os"
	"time"
)

// Config 定义项目的配置信息
type Config struct {
	ServerAddress string

	// JWT 相关配置
	SecretKey string

	// Redis 配置
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout time.Duration

	// 配置中心 & 服务发现（示例字段）
	ConfigCenterURL string
	NacosURL        string

	// Hystrix 熔断配置
	HystrixTimeout               int
	HystrixMaxConcurrentRequests int
	HystrixErrorPercentThreshold int
}

// LoadConfig 从环境变量加载配置，提供默认值
func LoadConfig() (*Config, error) {
	conf := &Config{
		ServerAddress:                getEnv("SERVER_ADDRESS", ":8080"),
		SecretKey:                    getEnv("JWT_SECRET", "your-secret-key"),
		RedisAddr:                    getEnv("REDIS_ADDR", "localhost:6379"),
		RedisMaxIdle:                 100,
		RedisMaxActive:               500,
		RedisIdleTimeout:             240 * time.Second,
		ConfigCenterURL:              getEnv("CONFIG_CENTER_URL", "http://config.example.com"),
		NacosURL:                     getEnv("NACOS_URL", "http://nacos.example.com"),
		HystrixTimeout:               1000,
		HystrixMaxConcurrentRequests: 100,
		HystrixErrorPercentThreshold: 50,
	}
	return conf, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}