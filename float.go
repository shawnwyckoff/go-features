package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	type item struct {
		x    float64
		s    string
		prec int
	}
	items := []item{
		item{x: 0.000123456789, s: "0.000123456789", prec: 12},
		item{x: 0.000123456789123456789000000, s: "0.000123456789123456789000000", prec: 21},
		item{x: 0.000123456789123456789123456789000000, s: "0.000123456789123456789123456789000000", prec: 30},
	}

	for _, v := range items {
		fmt.Println("Start:------")
		fmt.Println(v.s)
		fmt.Println("Float64")
		fmt.Println(strconv.FormatFloat(v.x, 'f', v.prec, 64))
		fmt.Println(strconv.FormatFloat(v.x, 'g', v.prec, 64))
		fmt.Println(fmt.Sprintf("%g", v.x))

		fmt.Println("Big Float")
		bf := big.NewFloat(v.x)
		fmt.Println(bf.String())
		fmt.Println(bf.Text('g', v.prec))
		fmt.Println(bf.Text('f', v.prec))
		fmt.Println(bf.Text('g', int(bf.Prec())))
		fmtstr := "%." + fmt.Sprintf("%d", v.prec) + "f"
		fmt.Println(fmt.Sprintf(fmtstr, bf))
		bf.SetString(v.s)
		fmt.Println(bf.String())
		fmt.Println(bf.Text('g', v.prec))
		fmt.Println(bf.Text('f', v.prec))
	}
}
