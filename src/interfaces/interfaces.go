package main

import "fmt"

type Shape interface {
	Area() float64
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	rect := Rectangle{2.2, 3.3}
	PrintArea(rect)
}
