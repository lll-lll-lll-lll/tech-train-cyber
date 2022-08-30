package api

import "github.com/gin-gonic/gin"

// InitRouter routerのパスを指定
func InitRouter() *gin.Engine {
	r := gin.Default()
	return r
}
