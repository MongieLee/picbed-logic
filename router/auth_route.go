package router

import (
	"github.com/gin-gonic/gin"
	"picbed/controllers"
)

func InitAuthRoute(e *gin.Engine) {
	authGroup := e.Group("/api/v1/auth")
	authController := controllers.AuthController{}
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
	}
}
