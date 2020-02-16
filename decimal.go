package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	a, _ := decimal.NewFromString("10000")
	b, _ := decimal.NewFromString("0.001436")
	fmt.Println(a.Div(b).Mul(b)) // 10000.0000000000000000000308
}
