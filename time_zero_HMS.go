package main

import (
	"fmt"
	"time"
)

func main() {
	tm1 := time.Date(2017, 07, 28, 00, 00, 00, 00, time.UTC)
	tm2 := time.Date(2017, 07, 28, 15, 48, 00, 00, time.UTC)
	if tm1.Before(tm2) {
		fmt.Printf("tm1 %s unix %d is before tm2 %s unix %d", tm1.String(), tm1.Unix(), tm2.String(), tm2.Unix())
	} else {
		fmt.Printf("tm2 %s unix %d is before tm1 %s unix %d", tm2.String(), tm2.Unix(), tm1.String(), tm1.Unix())
	}
}
