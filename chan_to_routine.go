package main

import (
	"fmt"
	"time"
)

func r1(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("input", i)
	}
}

func r2(ch chan int) {
	for i := 0; i < 100; i++ {
		n := <-ch
		fmt.Println("output", n)
	}
}

func main() {
	ch := make(chan int, 10)

	go r2(ch)
	time.Sleep(time.Second * 5)
	go r1(ch)

	for {
		time.Sleep(time.Second)
	}
}
