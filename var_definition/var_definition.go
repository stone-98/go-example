package main

import "fmt"

func main() {
	// 短变量声明，但是只能作用于函数内部
	a, aa := "", ""
	fmt.Println(a)
	fmt.Println(aa)

	// 依赖于字符串默认的初始化零值机制
	var b, bb string
	fmt.Println(b)
	fmt.Println(bb)

	var c, cc = "", ""
	fmt.Println(c)
	fmt.Println(cc)

	// 冗余声明类型
	var d, dd string = "", ""
	fmt.Println(d)
	fmt.Println(dd)
}
