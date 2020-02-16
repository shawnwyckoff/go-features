package main

import "fmt"

func sndChan(intCh chan int) {
	for i := 0; i < 10; i++ {
		intCh <- i
	}
}

func fun1_ok() {
	ch1 := make(chan int)
	go sndChan(ch1)
	for i := 0; i < 10; i++ {
		n := <-ch1
		fmt.Println(n)
	}
}

func sndChan2() chan int {
	intCh := make(chan int, 3)

	for i := 0; i < 10; i++ {
		intCh <- i
	}
	return intCh
}

// fixme: why?
func fun2_deadlock() {
	ch2 := sndChan2()
	go func() {
		for i := 0; i < 10; i++ {
			n := <-ch2
			fmt.Println(n)
		}
	}()
}

func main() {
	//fun1_ok()
	fun2_deadlock()
}
