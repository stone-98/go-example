package main

import "os"

func main() {
}

func open04(name string) error {
	// 这也是一种解决方案，将f的声明定义在if函数之外
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	f.Name()
	f.Close()
	return err
}
