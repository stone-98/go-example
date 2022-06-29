package _interface

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 接口规定
func interfaceStipulate() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	// 编译错误,time.Second未实现Writer接口中的方法
	// w = time.Second
	fmt.Println(w)

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	// 编译错误，bytes.Buffer未实现ReadWriteCloser
	//rwc = new(bytes.Buffer)
	fmt.Println(rwc)

	w = rwc
	// 编译错误rwc中有方法w中未实现
	//rwc = w
}

type IntSet struct {
	/****/
}

func (*IntSet) String() string {
	return "intSet"
}

func test() {
	os.Stdout.Write([]byte("hello"))
	os.Stdout.Close()

	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello"))
	//编译错误io.Writer没有Close方法
	//w.Close()
}
