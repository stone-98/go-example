package _func

func f(i, j, k int, s, t string) {
}

// 这个函数和上面那个函数意义一样
/*func f(i int, j int, k int, s string, t string) {
}*/

func add(x int, y int) int { return x + y }

// 类型相同可以简写
func sub(x, y int) (z int) { z = x - y; return }

// 不适用到得参数，可以用_代替
func first(x int, _ int) int { return x }

// 如果两个参数都没有使用到，都可以简写
func zero(int, int) int { return 0 }

// 没有消息体的方法，是其他编程语言实现的方法
//func Sin(x float64) float32
