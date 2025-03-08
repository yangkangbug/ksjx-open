package middleware

import (
	"net/http"
	"strings"

	"ksjx-open/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// CustomClaims 定义 JWT 的自定义 Claims 结构
type CustomClaims struct {
	UserInfo map[string]interface{} `json:"user"`
	jwt.StandardClaims
}

// JWTAuth 实现 JWT 认证中间件
func JWTAuth(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "缺少 Authorization 头"})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization 格式错误"})
			return
		}
		tokenStr := parts[1]
		token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.SecretKey), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的 Token"})
			return
		}
		claims, ok := token.Claims.(*CustomClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的 Token"})
			return
		}
		// 将用户信息保存到上下文中
		c.Set("user", claims.UserInfo)
		c.Next()
	}
}

// RBAC 实现权限校验中间件
func RBAC(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userData, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "没有权限"})
			return
		}
		user, ok := userData.(map[string]interface{})
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "权限信息错误"})
			return
		}
		roles, ok := user["roles"].([]interface{})
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "未分配角色"})
			return
		}
		hasRole := false
		for _, role := range roles {
			if role.(string) == requiredRole {
				hasRole = true
				break
			}
		}
		if !hasRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "权限不足"})
			return
		}
		c.Next()
	}
}
