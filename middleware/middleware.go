package middleware

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Auth 接口请求验证GitHub webhook
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		common, err := verifySignature(c)
		if err != nil || !common {
			c.Abort()
			return
		}
		fmt.Println("验证签名通过~")
		c.Next()
	}
}

// 验证签名
func verifySignature(c *gin.Context) (bool, error) {
	payloadBody, err := c.GetRawData()
	if err != nil {
		return false, err
	}

	// 获取请求头中的签名信息
	hSignature := c.GetHeader("X-Hub-Signature")

	// 计算Payload签名
	signature := hmacSha1(payloadBody)

	return hSignature == signature, nil
}

// hmac-sha1
func hmacSha1(payloadBody []byte) string {
	h := hmac.New(sha1.New, []byte("zruler"))
	h.Write(payloadBody)
	return "sha1=" + hex.EncodeToString(h.Sum(nil))
}
