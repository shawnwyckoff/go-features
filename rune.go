package main

import "fmt"

func main() {
	fmt.Println('b' - 'a')
	fmt.Println('2' - '1')
	var rst int = 0
	var c rune = 'c'
	var a rune = 'a'
	rst = int(c - a)
	fmt.Println(rst)
}
