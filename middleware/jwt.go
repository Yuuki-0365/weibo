package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"weibo/tool"
	"weibo/tool/e"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := tool.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if claims.ExpiresAt < time.Now().Unix() {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
