package main

import "fmt"

type S struct {
	Name string
}

func (s S) SetName(name string) {
	s.Name = name
}

func (s *S) SetName2(name string) {
	s.Name = name
}

func main() {
	s := S{Name: "no-name"}
	s.SetName("new-name")
	fmt.Println(s.Name)
	s.SetName2("new-name2")
	fmt.Println(s.Name)
}
