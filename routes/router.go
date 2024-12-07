package router

import (
	"TestDemo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	gin.SetMode(utils.AppMode)
	// Default默认加了日志和一个什么中间件
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
	return r
}
