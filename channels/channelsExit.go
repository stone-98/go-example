package main

import "fmt"

var done = make(chan struct{})

func main() {
	close(done)
	x, ok := <-done
	fmt.Println(ok)
	fmt.Println(x)
}
