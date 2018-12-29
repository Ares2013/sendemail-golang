package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/ws-sendemail-2/service"
	"net/http"
)

const (
	defaultPort = "8082"
)

// 中间件
func Corsmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(Corsmiddleware())
	router.GET("/ping", service.Ping)
	router.GET("/", service.Ping)

	// router group /v1
	v1 := router.Group("/v1")
	{
		v1.POST("/mails", service.SendEmailsV1)
		v1.GET("/mails", service.EmailsListV1)
		v1.GET("/mails/:email", service.EmailInfoV1)
		v1.POST("/resend/email/:email", service.ResendEmailsV1)
	}
	v2 := router.Group("/v2")
	{
		v2.POST("/mails", service.SendEmailsV1)
		v2.GET("/mails", service.EmailsListV1)
		v2.GET("/mails/:email", service.EmailInfoV1)
		v2.POST("/resend/email/:email", service.ResendEmailsV1)
	}
	router.Run(":" + defaultPort)
}
