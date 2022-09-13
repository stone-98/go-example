package multiplexing

import (
	"fmt"
	"testing"
)

func TestCountdown2(t *testing.T) {
	countdown2()
	<-abort
	fmt.Println("abort...")
}

func TestStdin(t *testing.T) {
	stdin()
}
