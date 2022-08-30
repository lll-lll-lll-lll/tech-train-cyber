package handler

import "github.com/gin-gonic/gin"

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func GetUser(c *gin.Context) {}

func CreateUser(c *gin.Context) {}

func UpdateUser(c *gin.Context) {}
