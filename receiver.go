package main

import "fmt"

// 总结
// 1、receiver为指针的函数才能修改其内部成员的内容，结构体则不可以
// 2、使用时则会隐式自动转换为结构体的指针
// 3、嵌套的结构体也可以通过指针receiver的函数修改内部成员的内容

type ExtendInfo struct {
	Address string
}

func (inf *ExtendInfo) SetAddress(address string) {
	inf.Address = address
}

type Student struct {
	Name     string
	Age      int
	MoreInfo ExtendInfo
}

func (s Student) SetName(name string) {
	s.Name = name
}

func (s *Student) SetAge(age int) {
	s.Age = age
}

func main() {
	var mingming Student = Student{Name: "mingming", Age: 10}
	mingming.MoreInfo.Address = "RenMin Road"

	mingming.SetName("tom")
	mingming.SetAge(11)
	mingming.MoreInfo.SetAddress("ShangHai Road")
	fmt.Println(mingming.Name, mingming.Age, mingming.MoreInfo.Address)
}
