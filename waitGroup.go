package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	fmt.Println("1")
	wg.Add(-1)
	wg.Wait()
	fmt.Println("-1")
	wg.Add(1)
	fmt.Println("1 again")
	wg.Add(-1)
	fmt.Println("-1 again")
	wg.Wait()
	wg.Wait()
	fmt.Println("wait")
	wg.Add(1)
	fmt.Println("1 twice")
}
