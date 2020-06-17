package main

import "fmt"

type People struct {
    Name string
}

func (p *People) String() string {
    return fmt.Sprintf("print: %v", p)
}

func main() {
    p := &People{}
    p.String()
}

/*
%v格式化字符串是本身会调用String()方法，这样用会导致无限递归，然后栈溢出。
*/
