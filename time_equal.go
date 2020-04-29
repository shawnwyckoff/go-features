package main

import (
	"fmt"
	"time"
)

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	tm1 := time.Now().In(loc)
	tm2 := tm1.In(time.UTC)
	if tm1.Equal(tm2) {
		fmt.Println("Time equal judgment is NOT affected by time zone")
	}
}
