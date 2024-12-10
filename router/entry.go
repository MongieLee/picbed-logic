package router

import (
	"github.com/gin-gonic/gin"
	"picbed/controllers"
	"picbed/middlewares"
)

func initMiddleware(e *gin.Engine) {
	e.Use(middlewares.CorsHandler)
}

func InitRoutes(e *gin.Engine) {
	e.NoRoute(func(c *gin.Context) {
		controllers.ResponseWithFail(c, controllers.CodeResourceNotFound)
	})
	initMiddleware(e)
	InitAuthRoute(e)
	InitUserRoute(e)
}
