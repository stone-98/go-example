package main

import "os"

func main() {

}

func open05(name string) error {
	// 这也是一种解决方案，将f的使用写在else里面，因为在go语言中，else的作用域算是在if内的
	if f, err := os.Open(name); err != nil {
		return err
	} else {
		f.Name()
		f.Close()
	}
	return nil
}
