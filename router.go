package main

import (
	"gin-gorm/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/auth/register", controller.UserRegister)
	r.POST("/api/auth/login", controller.Login)
	return r
}
