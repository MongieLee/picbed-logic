package main

import (
	"github.com/gin-gonic/gin"
	"picbed/global"
	"picbed/router"
	"picbed/utils"
)

func main() {
	global.InitGlobalEnv()
	r := gin.Default()
	router.InitRoutes(r)
	utils.InitZeroLog()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": struct{ name string }{name: "1231231"},
		})
	})
	r.Run(":9999")
}
