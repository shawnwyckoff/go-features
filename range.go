package main

import "fmt"

// range一个数组时，i的值是提前排列好的，不可以在循环内改变

func main() {
	ss := []string{"a", "b", "c", "d", "e", "f", "g"}

	for i := range ss {
		fmt.Println(ss[i])

		if ss[i] == "b" {
			i--
		}
	}
}
