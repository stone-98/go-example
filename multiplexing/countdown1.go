package multiplexing

import (
	"fmt"
	"time"
)

func countdown1() {
	fmt.Println("Commencing countdown.")
	// time.Tick函数，在countdown函数返回时，它会停止从time.Tick中接收事件，但是ticker这个goroutine还依然存活，继续的像channel中发送值
	// 然而这个时候已经没有其他的goroutine会从该channel中接收了，这就是goroutine的泄露
	// time.Tick函数挺方便，但是只有在程序的整个生命周期都需要时才用它
	// 否则我们应该使用time.NewTicker()
	// ticker.Stop() 不用了就停止。
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		time := <-tick
		fmt.Println(time)
	}
}
