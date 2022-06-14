//go:build ignore
// +build ignore

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	// 3.监听端口（如果不指定默认为8080端口）
	r.Run(":8080")
}
