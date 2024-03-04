package main

import (
	"AutoDeploy/server"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.GET("/deploy", server.AutoDeploy)
	err := engine.Run(":8989")
	if err != nil {
		return
	}
	fmt.Println("自动部署脚本运行成功··········")
}
