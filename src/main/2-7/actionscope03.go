package main

import "os"

func main() {

}

func open() error {
	if f, err := os.Open("123"); err != nil {
		return err
	}
	// 变量f的作用域只在if语句内，因此后面的语句将无法引入它，这将导致编译错误。
	f.ReadByte()
	f.Close()
	return nil
}
