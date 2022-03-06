package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func main() {
	// var声明数组
	var a [3]int
	// 数组使用下标来访问,没有初始化的元素都会被初始化为元素类型对应的零值
	fmt.Println(a[0])
	// len()返回数组的长度
	fmt.Println(a[len(a)-1])
	// 使用数组字面值来初始化数组
	var q [3]int = [3]int{1, 2, 3}
	fmt.Println(q)
	// 声明r数组的容量为3，但是只给其中两个元素赋值
	var r [3]int = [3]int{1, 2}
	// 没赋值的还是会默认被初始化为零值
	fmt.Println(r[2])
	// 使用...来省略数组的长度，则f的数组长度是由初始化值的个数来计算
	f := [...]int{1, 2, 3}
	fmt.Printf("%T\n", f)
	// 数组的长度是数组的一个组成部分，所以长度为3和长度为4的数组，他们是不同的数组类型，数组的长度必须是一个常量表达式，所以需要在编译期间确定
	t := [3]int{1, 2, 3}
	//t = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
	fmt.Println(t)
	// TODO 没太看明白，因为没有看常量的使用
	// 上面的形式是直接提供顺序初始化值序列，但是也可以指定一个索引和对应值列表的方式初始化
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"
	// 定义了一个含有100个元素的数组r，最后一个元素被初始化为-1，其它元素都是用0初始化，这种方式没有被初始化的元素还是会给零值
	j := [...]int{99: -1}
	fmt.Println(j)
	// 数组之间的比较，只有数组中的每个元素相等，那么数组才会相等
	aa := [3]int{1, 2, 3}
	var bb = [3]int{1, 2, 3}
	var cc [2]int = [2]int{1, 2}
	fmt.Println(cc)
	fmt.Println(aa == bb)
	//fmt.Println(bb == cc) compile error 数组的类型不相等，不能进行比较
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("y"))
	// %x：格式化 16 进制表示的数字
	// %t：格式化布尔型
	// %T：打印某个类型的完整说明
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	tt := [32]byte{1, 2, 3}
	invalidZero(tt)
	fmt.Println(tt)
	zero(&tt)
	fmt.Println(tt)
}

// 当调用一个函数的时候，函数的每个调用参数将会被赋值给函数内部的参数变量，所以函数参数变量接收的是一个复制的副本，并不是原始调用的变量。
// 所以这个初始化是无效的
func invalidZero(prt [32]byte) {
	prt = [32]byte{}
}

// 使用指针才能有效的重新初始化为0
func zero(prt *[32]byte) {
	*prt = [32]byte{}
}
