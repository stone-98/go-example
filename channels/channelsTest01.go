package main

import (
	"fmt"
	"math/rand"
	"time"
)

var channel = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v \n", value)
	channel <- value
}

func main() {
	defer close(channel)
	go send()
	fmt.Println("wait...")
	value := <-channel
	fmt.Printf("receive : %v \n", value)
	fmt.Println("end...")
}
