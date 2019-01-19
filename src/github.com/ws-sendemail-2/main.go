package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/ws-sendemail-2/service"
	"net/http"
	"os"
		"fmt"
	"time"
	)

const (
	defaultPort = "8082"
)
// golang新版本的应该
func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

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
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	// _time := time.Now().String()
	// PathExist(_time)
	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(Corsmiddleware())
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
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
