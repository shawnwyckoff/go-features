package main

import "fmt"

var receiverSummary = `
1、receiver为指针的函数才能修改其内部成员的内容，结构体则不可以
2、使用时则会隐式自动转换为结构体的指针
3、嵌套的结构体也可以通过指针receiver的函数修改内部成员的内容

4、指针Receiver的函数，Receiver就是自己，可以操作内部成员
5、非指针Receiver的函数，Receiver是个拷贝，不可以操作内部成员
`

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
	mm := Student{Name: "mingming", Age: 10}
	mm.MoreInfo.Address = "RenMin Road"

	mm.SetName("tom")
	mm.SetAge(11)
	mm.MoreInfo.SetAddress("ShangHai Road")
	fmt.Println(mm.Name, mm.Age, mm.MoreInfo.Address)
}
