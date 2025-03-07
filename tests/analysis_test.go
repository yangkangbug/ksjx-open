package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"ksjx-open/config"
	"ksjx-open/middleware"
	"ksjx-open/router"
	"ksjx-open/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// 模拟回源查询函数
func fetchFromSource(key string) ([]byte, error) {
	return []byte("source_data"), nil
}

func TestCache_GetWithCache(t *testing.T) {
	// 初始化 Redis 连接池（可指向本地测试 Redis）
	service.InitRedisPool("localhost:6379")

	// 测试本地缓存
	key := "test_key"
	// 初次查询，回源查询数据存入缓存
	data, err := service.GetWithCache(key, fetchFromSource)
	if err != nil || string(data) != "source_data" {
		t.Errorf("首次查询失败，应返回 source_data, 得到 %v, 错误: %v", data, err)
	}
	// 第二次查询应直接从缓存返回
	data, err = service.GetWithCache(key, fetchFromSource)
	if err != nil || string(data) != "source_data" {
		t.Errorf("二次查询未从缓存返回，得到 %v", data)
	}
}

func TestJWTAuth(t *testing.T) {
	conf, _ := config.LoadConfig()
	// 构造一个测试 Token
	claims := &middleware.CustomClaims{
		UserInfo: map[string]interface{}{"id": "user_123", "roles": []interface{}{"admin"}},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(conf.SecretKey))

	// 构造测试请求并添加 Authorization 头
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.JWTAuth(conf))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	})

	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("JWTAuth 测试失败，状态码应为 200，实际为 %d", w.Code)
	}
}

func TestRouter_Health(t *testing.T) {
	conf, _ := config.LoadConfig()
	r := router.InitRouter(conf)
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "OK") {
		t.Errorf("健康检查接口返回错误，状态码: %d, body: %s", w.Code, w.Body.String())
	}
}
