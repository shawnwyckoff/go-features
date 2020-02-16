package main

import "fmt"

// defer虽然是在return之前执行的，但是如果return后面接的是函数，那么这个函数还是先于defer执行的

func sub() error {
	fmt.Println("sub")
	return nil
}

func fun() error {
	fmt.Println("main begin")

	defer fmt.Println("main defer")
	return sub()
}

func main() {
	fun()
}
