package main

import "fmt"

func main() {
	// 声明x
	x := "hello!"
	// 声明了i，使用了外部的x
	for i := 0; i < len(x); i++ {
		// 重新声明了一个x
		x := x[i]
		// 使用外面一层的x，编译器会从里往外找
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}
}
