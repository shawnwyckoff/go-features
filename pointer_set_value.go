package main

import (
	"fmt"
	"time"
)

var (
	ZeroTime  time.Time = time.Time{} // 0001-01-01 00:00:00 +0000 UTC
	ZeroTime2 time.Time = time.Time{} // 0001-01-01 00:00:00 +0000 UTC
)

func main() {
	var ptm *time.Time = nil
	ptm = &ZeroTime
	*ptm = time.Now()
	if ZeroTime.Equal(ZeroTime2) == false {
		fmt.Println(ZeroTime)
		fmt.Println("please be careful to use *Pointer=Variable, this may change original variable value")
	}
}
