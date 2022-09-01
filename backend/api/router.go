package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tech-train-cyber/api/handler"
)

// InitRouter routerのパスを指定
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authorizedUserAPI := r.Group("/api", AuthRequired())
	{
		authorizedUserAPI.GET("/user/get", handler.GetUser)
		authorizedUserAPI.POST("/user/create", handler.CreateUser)
		authorizedUserAPI.POST("/user/update", handler.UpdateUser)
	}

	return r
}

//
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
