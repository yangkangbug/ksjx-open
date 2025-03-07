package service

import (
	"math/rand"
	"time"
)

// Select 随机负载均衡，从实例列表中随机选择一个服务地址
func Select(instances []string) string {
	if len(instances) == 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	return instances[rand.Intn(len(instances))]
}
