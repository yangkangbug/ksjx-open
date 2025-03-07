package service

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// LocalCache 定义简单的内存缓存
var LocalCache = NewSimpleCache()

// SimpleCache 是一个简单的内存缓存实现
type SimpleCache struct {
	items map[string][]byte
}

// NewSimpleCache 初始化一个简单缓存
func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		items: make(map[string][]byte),
	}
}

// Get 获取缓存数据
func (c *SimpleCache) Get(key string) (interface{}, bool) {
	val, ok := c.items[key]
	return val, ok
}

// Set 存储数据到缓存
func (c *SimpleCache) Set(key string, value []byte) {
	c.items[key] = value
}

// RedisPool 全局 Redis 连接池
var RedisPool *redis.Pool

// InitRedisPool 初始化 Redis 连接池
func InitRedisPool(redisAddr string) {
	RedisPool = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   500,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisAddr)
		},
	}
}

// GetWithCache 多级缓存：先查内存，再查 Redis，最后回源查询
func GetWithCache(key string, fetchFromSource func(string) ([]byte, error)) ([]byte, error) {
	// 内存缓存查询
	if val, ok := LocalCache.Get(key); ok {
		return val.([]byte), nil
	}
	// Redis 缓存查询
	conn := RedisPool.Get()
	defer conn.Close()
	val, err := redis.Bytes(conn.Do("GET", key))
	if err == nil && val != nil {
		LocalCache.Set(key, val)
		return val, nil
	}
	// 回源查询
	data, err := fetchFromSource(key)
	if err == nil {
		LocalCache.Set(key, data)
		conn.Do("SETEX", key, 3600, data)
	}
	return data, err
}
