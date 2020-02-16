package main

import "fmt"

var selectCaseSummary = `
每个case都必须是一个执行 <- 运算的channel通信
所有channel表达式都会被求值，所有被发送的表达式都会被求值
case和default的路径优先级：case优先级大于default。如果所有case都阻塞，且有default子句，则执行default。如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
case和case间的路径优先级： case之间优先级相同，如果有多个case都不阻塞，select会随机公平地选出一个执行，其他不会执行。
`

// 用select检测chan是否已满，或者理解为，尽量写入chan，满了就放弃
func main() {
	ch := make(chan int, 1)
	ch <- 1

	select {
	case ch <- 2:
	default:
		fmt.Println("channel is full !")
	}
}