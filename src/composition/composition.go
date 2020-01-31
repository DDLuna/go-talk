package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) sayHi() {
	fmt.Println("hi")
}

type Student struct {
	Person
	AvgScore float64
}

func main() {
	student := Student{}
	student.sayHi()
}
