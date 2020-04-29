package main

import (
	"fmt"
	"time"
)

type MyDate time.Time

func NewMyDate(year, month, day int, tz time.Location) (MyDate, error) {
	return MyDate(time.Date(year, time.Month(month), day, 0, 0, 0, 0, &tz)), nil
}

func (md MyDate) ToTime() time.Time {
	return time.Time(md)
}

func main() {
	timeMap := make(map[time.Time]string)

	md1, _ := NewMyDate(2019, 1, 1, *time.UTC)
	md2, _ := NewMyDate(2019, 1, 1, *time.UTC)

	timeMap[md1.ToTime()] = "1"
	timeMap[md2.ToTime()] = "2"

	for k, v := range timeMap {
		fmt.Println(k, v)
	}
}
