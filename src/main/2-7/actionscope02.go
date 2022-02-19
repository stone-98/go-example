package main

import (
	"log"
	"os"
)

var cwd2 string

// 这是一种解决方案，不使用":="进行初始化
func main() {
	var err error
	cwd2, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
