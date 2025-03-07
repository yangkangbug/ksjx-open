package router

import (
	"ksjx-open/client"

	"github.com/gin-gonic/gin"
)

// RegisterInventoryRoute 注册查询库存的 HTTP 接口
func RegisterInventoryRoute(r *gin.Engine) {
	// 定义 URL 路径，参数 productId
	r.GET("/inventory/:productId", queryInventory)
}

// queryInventory 处理查询库存接口请求
func queryInventory(c *gin.Context) {
	productId := c.Param("productId")
	i, err := client.QueryInventory(productId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"stock": i})
}
