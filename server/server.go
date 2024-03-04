package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/exec"
)

func AutoDeploy(c *gin.Context) {
	err := execute()
	if err != nil {
		return
	}
	c.JSON(200, "自动部署成功~~~")
}

func execute() error {
	cmd := exec.Command("sh", "/Users/xwx/go/src/AutoDeploy/server/exec.sh")
	// 打开或创建日志文件
	logFile, err := os.OpenFile("log/exec.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
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
