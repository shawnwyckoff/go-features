package main

import (
	"fmt"
	"time"
)

func main()  {
	tm := time.Date(2019,9,1,0,0,0,0, time.Local)
	fmt.Println(tm.Round(time.Hour * 24)) // 2019-09-01 08:00:00 +0800 CST
}
