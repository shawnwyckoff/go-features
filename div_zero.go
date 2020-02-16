package main

import (
	"fmt"
	"math"
)

func main() {
	zero := 0

	dived := 1.2 / float64(zero)
	fmt.Println(dived)

	if math.IsInf(dived, 0) {
		fmt.Println("dived is infinity")
	}
}
