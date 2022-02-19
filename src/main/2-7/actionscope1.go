package main

import "fmt"

func f() {

}

var g = "g"

func main() {
	f := "F"
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h) // 编译错误
}
