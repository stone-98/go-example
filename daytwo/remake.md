- os包提供了一些与操作系统交互的函数和变量，程序的命令行参数可以从os包的Args参数中获取，os.Args是一个字符串切片，第一个元素os.Args[0]是命令本身的名字，其他的元素则是程序启动传给它的参数。
- import的写法两种：

```go
package main

import (
	"fmt"
	"os"
)

//or
//import "fmt"
//import "os"

func main() {
	fmt.Println(os.Args)
}
```

- go中的for循环

```go
package main

import (
	"fmt"
	"os"
)

//for init; condition; post{
//
//}
//
//or

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
}
```

init、condition、post都可以省略，
如果全部省略则是一个无线循环
如果省略init和post分号也可以省略

- 初始化变量的方式
```
    // 只能在函数内部声明，不能用于包变量
    s := ""
    // 依赖于字符串的默认初始化零值机制，被初始化为""。
    var s string
    // 用得很少，除非同时声明多个变量。
    var s = ""
    // 形式显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了。
    var s string = ""
```


- strings.Join()合并数组

- todo 测试