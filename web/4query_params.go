//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("user", func(context *gin.Context) {
		name := context.DefaultQuery("name", "阿魁")
		context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	router.Run()
}
