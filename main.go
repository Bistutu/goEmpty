package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"GoEmpty/utils"
	"GoEmpty/utils/httputil"
	"GoEmpty/utils/loginutil"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery()) // 使用 Recovery 中间件
	// 自定义日志格式
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s \" %d %s \"%s\" \"%s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.GET("/login", func(context *gin.Context) {
		httputil.RemoveAllCookie() // 清空所有 cookie
		username := context.Query("username")
		password := context.Query("password")
		if username == "" || password == "" || len(username) < 10 {
			log.Printf("a bad request: %v\n", username)
			context.String(http.StatusBadRequest, "")
			return
		}
		cookies, err := loginutil.GetEmptyCookie(username, password)

		if err != nil {
			log.Fatalf("login fail: %v", err)
			context.String(http.StatusBadRequest, "")
			return
		}
		context.String(http.StatusOK, utils.CookieToPlainStr(cookies))
	})
	router.Run(":9999")
}
