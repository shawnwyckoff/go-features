package main

import (
	"fmt"
	"time"
)

func run() {
	go func() {
		for i := 0; i < 1024000; i++ {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}()
}

func main() {
	run()

	fmt.Println("run over")

	for {
		time.Sleep(time.Second)
	}
}
