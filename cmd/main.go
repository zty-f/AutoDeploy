package main

import (
	"AutoDeploy/middleware"
	"AutoDeploy/server"
	"fmt"
	"github.com/gin-gonic/gin"
)

// curl http://localhost:8989/ping
// curl -X POST http://localhost:8989/deploy

func main() {
	engine := gin.New()
	engine.Use(middleware.TimeoutMiddleware())
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "Welcome to zty project!")
	})
	engine.POST("/deploy", middleware.Auth(), server.AutoDeploy)
	err := engine.Run(":8989")
	if err != nil {
		return
	}
	fmt.Println("自动部署脚本运行成功··········")
}
