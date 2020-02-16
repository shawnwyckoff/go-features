package main

import (
	"fmt"
	"reflect"
)

// shows how to check whether 2 maps equal or not

func main() {
	m1 := make(map[string]string)
	m2 := make(map[string]string)

	m1["1"] = "abc"
	m1["2"] = "一二三"

	m2["2"] = "一二三"
	m2["1"] = "abc"

	fmt.Println(reflect.DeepEqual(m1, m2))
}
