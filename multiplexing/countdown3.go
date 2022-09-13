package multiplexing

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	// create abort channel
	abort := make(chan struct{})
	go func() {
		// 需要调整监控键盘的输入
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("await finish.")
	case <-abort:
		fmt.Println("Launch aborted!")
	default:
		fmt.Println("while")
	}
	doLaunch()
}
func doLaunch() {
	fmt.Println("do launch.")
}
