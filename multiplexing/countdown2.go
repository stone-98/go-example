package multiplexing

import (
	"fmt"
	"os"
)

var abort = make(chan struct{})

func countdown2() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
}

func stdin() {
	var buffer [512]byte
	n, err := os.Stdin.Read(buffer[:])
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	fmt.Println("count:", n, ",msg:", string(buffer[:]))
}
