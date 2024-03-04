package main

import (
	"AutoDeploy/server"
	"fmt"
	"github.com/gin-gonic/gin"
)

// curl http://localhost:8989/ping
// curl http://localhost:8989/deploy

func main() {
	engine := gin.New()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "Welcome to zty project!")
	})
	engine.GET("/deploy", server.AutoDeploy)
	err := engine.Run(":8989")
	if err != nil {
		return
	}
	fmt.Println("自动部署脚本运行成功··········")
}
