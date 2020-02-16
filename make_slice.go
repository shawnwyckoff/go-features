package main

import "fmt"

/**
make陷阱
*/

func main() {
	s := make([]int, 3)
	s = append(s, 1, 2, 3)
	fmt.Println(s) // [0 0 0 1 2 3]
}
