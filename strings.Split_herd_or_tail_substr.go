package main

import (
	"fmt"
	"strings"
)

func main()  {
	ss1 := strings.Split("head.ABC.tail", "head")
	ss2 := strings.Split("head.ABC.tail", "tail")
	if len(ss1) == 2 && len(ss2) == 2 {
		fmt.Println("if split head or tail substr, there will be an empty item found in result array")
	}
}
