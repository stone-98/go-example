package main

import "fmt"
// Go语言主要有四种类型的声明语句：var、const、type和func
//包一级范围声明语句声明的，包一级声明语句声明的名字可在整个包对应的每个源文件中访问
const name = "stone"

// n、s是在main函数内部声明的声明语句声明的，局部声明的名字就只能在函数内部很小的范围被访问
func main() {
	var n = name
	var s = name + "98"
	fmt.Printf("%s , %s", n, s)
}
