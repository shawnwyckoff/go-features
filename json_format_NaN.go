package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type JFN struct {
	N float64
}

func main() {
	jfn := JFN{N: math.NaN()}
	b, err := json.Marshal(jfn)
	if err != nil {
		fmt.Println("can't Marshal float64 which is math.NaN()")
		panic(err)
	}
}
