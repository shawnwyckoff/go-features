package main

import "fmt"

// Animal interface
type Animal interface {
	Eat()
	Born() Animal
}

// Cat type
type Cat struct {
}

// If returns Cat, compiler print errors, you should use Animal
func (c Cat) Born() Cat {
	fmt.Println("Cat born")
	return Cat{}
}

func (c Cat) Eat() {
	fmt.Println("Cat eat")
}

// Function use Animal interface
func BornBaby(animal Animal) {
	animal.Born()
}

func main() {
	cat := Cat{}

	BornBaby(cat)
}
