package router

import (
	"github.com/gin-gonic/gin"
	"picbed/controllers"
	"picbed/middlewares"
)

func InitCommonRoute(e *gin.Engine) {
	fileController := controllers.FileController{}
	group := e.Group("/api/v1/common")
	group.Use(middlewares.AuthMiddleware)
	{
		group.POST("/uploadFile", fileController.UploadFile)
	}
}
