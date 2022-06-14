//go:build ignore
// +build ignore

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	router := gin.Default()
	router.GET("/user/:name/*active", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("active")
		action = strings.Trim(action, "/")
		context.String(http.StatusOK, name+action)
	})
	router.Run()
}
