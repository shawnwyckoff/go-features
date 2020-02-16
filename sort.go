package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type StudentSet struct {
	Items []Student
}

func (ss StudentSet) Len() int {
	return len(ss.Items)
}

func (ss StudentSet) Swap(i, j int) {
	ss.Items[i], ss.Items[j] = ss.Items[j], ss.Items[i]
}

func (ss StudentSet) Less(i, j int) bool {
	return ss.Items[i].Age < ss.Items[j].Age
}

func main() {
	ss := StudentSet{}
	ss.Items = append(ss.Items, Student{Name: "Tom", Age: 10})
	ss.Items = append(ss.Items, Student{Name: "Jerry", Age: 12})
	ss.Items = append(ss.Items, Student{Name: "David", Age: 11})
	ss.Items = append(ss.Items, Student{Name: "Elon", Age: 9})
	ss.Items = append(ss.Items, Student{Name: "Satoshi", Age: 13})
	sort.Sort(ss)

	fmt.Println(ss)
}
