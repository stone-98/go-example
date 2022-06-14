//go:build ignore
// +build ignore

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word!")
	})
	r.POST("/post", func(context *gin.Context) {
		// no handler
	})
	r.PUT("/put", func(context *gin.Context) {
		// no handler
	})
	// 默认为8080端口
	r.Run()
}
