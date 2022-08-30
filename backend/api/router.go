package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-train-cyber/api/handler"
)

// InitRouter routerのパスを指定
func InitRouter() *gin.Engine {
	r := gin.Default()
	userapi := r.Group("/api")
	{
		userapi.GET("/user/get", handler.GetUser)
		userapi.POST("/user/create", handler.CreateUser)
		userapi.POST("/user/update", handler.UpdateUser)
	}
	return r
}
