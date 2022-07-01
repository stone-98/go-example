package _interface

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Value() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	w = nil
	fmt.Println(w)
}

func Incommensurability() {
	// 切面不能比较
	// TODO 但是为什么能这样写呀？
	var x interface{} = []int{1, 2, 3}
	fmt.Println(x == x)
}

func PrintInterfaceType() {
	var w io.Writer
	fmt.Println("%T\n", w)
	w = os.Stdout
	fmt.Println("%T\n", w)
	w = new(bytes.Buffer)
	fmt.Println("%T\n", w)
}

const debug = false

func InterfaceFeatures() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	doWrite(buf)
}

func doWrite(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
