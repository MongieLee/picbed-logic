package router

import (
	"github.com/gin-gonic/gin"
	"picbed/controllers"
	"picbed/middlewares"
)

func InitUserRoute(e *gin.Engine) {
	userController := controllers.UserController{}
	group := e.Group("/api/v1/user")
	group.Use(middlewares.AuthMiddleware)
	{
		group.POST("/create", userController.CreateUser)
		group.POST("/update", userController.UpdateUser)
		group.POST("/delete", userController.DeleteUser)
		group.GET("/getById", userController.GetUserById)
		group.GET("/getList", userController.GetUsers)
	}
}
