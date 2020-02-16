package main

import "fmt"

type myType struct {
	a int
}

// don't init member when new struct
func main() {
	mt := new(myType{a: 0}) // this will report error: myType literal is not a type
	fmt.Println(mt)
}
