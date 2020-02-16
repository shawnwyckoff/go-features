package main

/*
总结：
slice浅拷贝的操作：slice的[]切片
slice深拷贝的操作：append, make有效长度后copy
copy(dst, src)时如果dst的初始化没有用make([]TypeSide, size)而是var dst []TypeSide，那么copy无效，因为dst的长度是0，所以拷贝数据长度也为0
以上特性对基本类型的slice有效，对自定义结构体等类型的slice同样有效, 嵌入在结构体中的slice同样有效，甚至整个结构体拷贝之后，内部的slice也是如此

切片下标的含义：
b := a[x:y], y和x一样都是下标不是length，区别就是x包括在拷贝范围，而y不
*/

import (
	"fmt"
	"github.com/shawnwyckoff/go-utils/dsa/bytes2"
	"github.com/shawnwyckoff/go-utils/dsa/bytez"
)

func main() {

	a := make([]int, 5, 10)
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 5
	//a  len: 5 cap: 10
	fmt.Println("a: ", "len:", len(a), " cap:", cap(a))

	b := a[:3]
	//b  len: 3 cap: 10
	fmt.Println("b: ", "len:", len(b), " cap:", cap(b))

	//当从左边界有截断时 会改变新切片容量大小
	c := a[2:4]
	c[0] = 6
	//c  len: 2 cap: 8
	fmt.Println("c: ", "len:", len(c), " cap:", cap(c))

	//左边界默认0 最小为0 | 右边界默认slice的长度 最大为slice的容量
	d := a[:cap(a)]
	d[0] = 7
	//d  len: 10 cap: 10
	fmt.Println("d: ", "len:", len(d), " cap:", cap(d))

	fmt.Println(a)

	// 结构体slice一样默认是浅拷贝
	type S struct {
		N int
	}

	type E struct {
		Ss []S
	}

	s1 := make([]S, 3)
	s1[0].N = 0
	s1[1].N = 1
	s1[2].N = 2

	s2 := s1[:]

	var s3 []S
	copy(s3, s1)

	s4 := make([]S, len(s1))
	s4 = s1[:]

	s5 := s1

	var s7 []S
	s7 = append(s7, s1...)

	s8 := make([]S, len(s1))
	copy(s8, s1)

	e1 := E{Ss: s1}
	e2 := e1
	e3itf, err := bytez.DeepCopyEx(e1)
	fmt.Println(err)

	s2[0].N = 1

	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("s3: ", s3)
	fmt.Println("s4: ", s4)
	fmt.Println("s5: ", s5)
	//fmt.Println("s6: ", s6)
	fmt.Println("s7: ", s7)
	fmt.Println("s8: ", s8)

	fmt.Println("e1: ", e1)
	fmt.Println("e2: ", e2)
	fmt.Println("e3: ", e3itf.(E))

	sx := s1[2:3]
	fmt.Println(sx)

}
