package main

import (
	"fmt"
	"os"
)

func main() {
	var params string
	for i, s := range os.Args {
		fmt.Printf("\nindex:[%v],param:[%v]\n", i, s)
		params += s
	}
	fmt.Printf("all:[%v]", params)
}
