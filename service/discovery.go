package service

import (
	"fmt"
)

// GetService 模拟从 Nacos 中获取服务实例列表
func GetService(serviceName string) []string {
	// 此处为伪实现，实际中需要调用 Nacos SDK 获取服务列表
	fmt.Printf("从 Nacos 获取服务：%s\n", serviceName)
	return []string{
		"http://127.0.0.1:8081",
		"http://127.0.0.1:8082",
	}
}
