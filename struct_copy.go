package main

import (
	"fmt"
)

/**
结构体是值传递，但是其中的成员如果是slice等引用传递类型，那么该成员依然只能是引用传递*/

type (
	People struct {
		Name string
		Tags []string
	}
)

func main() {
	a := People{Name: "Lucy", Tags: []string{"1", "2", "3"}}
	b := a

	b.Tags[0] = "tag1"
	fmt.Println(a)
	fmt.Println(b)
}
