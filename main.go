package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"os"
	"picbed/global"
	"picbed/router"
	"picbed/utils"
)

func main() {
	command := &cobra.Command{
		Use:   "run",
		Short: "Run picbed logic server!",
		Run: func(cmd *cobra.Command, args []string) {
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
		},
	}
	command.Flags().StringVarP(&global.ViperConfigFile, "config", "c", "", "config file path")
	command.MarkFlagRequired("config")
	rootCommand := &cobra.Command{}
	rootCommand.AddCommand(command)
	if err := rootCommand.Execute(); err != nil {
		fmt.Println("启动遇到错误：", err.Error())
		os.Exit(0)
	}
}
