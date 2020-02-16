package main

import (
	"encoding/json"
	"fmt"
)

type Demo struct {
	Float *float64
}

func main() {
	flt := 123.456
	demo := Demo{Float: &flt}
	if buf, err := json.Marshal(&demo); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(buf))
	}

	demo2 := Demo{}
	if buf, err := json.Marshal(&demo2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(buf))
	}
}
