package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{"hello": "world"},
	)
}

func CreateUser(c *gin.Context) {}

func UpdateUser(c *gin.Context) {}
