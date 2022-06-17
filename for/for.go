package main

import "fmt"

func main() {
	var str = "hello"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	a := 0
	for true {
		if a > 1 {
			break
		}
		fmt.Println("while true...")
		a++
	}

	b := 0
	for {
		if b > 1 {
			break
		}
		fmt.Println("while...")
		b++
	}

	for _, c := range str {
		fmt.Println(c)
	}

	for index := range str {
		fmt.Println(index)
	}

	for index, c := range str {
		fmt.Println(c)
		fmt.Println(index)
	}
}
