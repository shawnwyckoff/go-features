package main

import "fmt"

var deferSummary = `
1、defer虽然是在return之前执行的，但是如果return后面接的是函数，那么这个函数还是先于defer执行的
2、defer会在return指令之前执行，而go中的“return **”不是原子操作，其返回值放在栈而不是寄存器（比如C），所以“return **”执行时会被拆解为“$返回变量 = ** return”，所以defer后面的语句会在“$返回变量 = **“ 和 ”return”之间执行
`

// 返回值应该是1
func deferF1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

// 返回值应该是1
func deferF2() (r int) {
	defer func(r int) { // r是通过传参拷贝的，所以不会改变f()中r的值
		r = r + 5
	}(r)
	return 1
}

func deferF3() {
	fmt.Println("F3 begin")
	defer fmt.Println("F3 defer")
}

func main() {
	deferF3()
}
