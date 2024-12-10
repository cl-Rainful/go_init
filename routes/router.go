package router

import (
	v1 "TestDemo/api/v1"
	"TestDemo/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	gin.SetMode(utils.AppMode)
	// Default默认加了日志和一个什么中间件
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		// 用户模块
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.PUT("user/:id", v1.EditUser)
		routerV1.DELETE("user/:id", v1.DelUser)
		// 分类模块

		// 文章模块
	}
	return r
}
