package main

import (
	"log"
	"os"
)

var cwd1 string

func main() {
	cwd1, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	// 编译器会判断错误，虽然包级cwd变量没有被使用，但是还是可以编译通过
	log.Printf("Working directory = %s", cwd1)
}
