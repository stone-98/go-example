package _interface

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

// 测试接口规定
func TestInterfaceStipulate(t *testing.T) {
	interfaceStipulate()
}

// 测试指针方法
func TestPointerMethod(t *testing.T) {
	// 编译错误，不能在一个不能寻址的IntSet值上调用String()
	//var _ = IntSet{}.String()
	var intSet IntSet
	var _ = intSet.String()

	// *IntSet实现了String方法，所以下行可以成功编译
	var _ fmt.Stringer = &intSet
	// 编译错误
	//var _ fmt.Stringer = intSet
}

func TestNilInterface(t *testing.T) {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	fmt.Println(any)
}

func TestNewAssert(t *testing.T) {
	// 断言*bytes.Buffer实现了io.Writer接口
	var w io.Writer = new(bytes.Buffer)
	fmt.Println(w)
	var _ io.Writer = (*bytes.Buffer)(nil)
}
