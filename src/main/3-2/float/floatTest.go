package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x= %g e^x = %8.3f  e^x = %e \n", x, math.Exp(float64(x)), math.Exp(float64(x)), math.Exp(float64(x)))
	}

	nan := math.NaN()
	fmt.Println(math.IsNaN(nan))
	fmt.Println(nan == nan, nan < nan, nan > nan)
}
