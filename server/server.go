package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/exec"
)

func AutoDeploy(c *gin.Context) {
	action := c.GetHeader("X-GitHub-Event")
	if action == "ping" {
		c.JSON(200, "测试webhook连接成功~")
	}
	err := execute()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "自动部署成功~~~")
}

func execute() error {
	cmd := exec.Command("sh", "exec.sh")
	// 打开或创建日志文件
	logFile, err := os.OpenFile("../log/deploy_exec.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()
	cmd.Stdout = logFile
	cmd.Stderr = logFile

	if err = cmd.Run(); err != nil {
		log.Printf("Error starting command: %s", err.Error())
		return err
	}

	return nil
}
