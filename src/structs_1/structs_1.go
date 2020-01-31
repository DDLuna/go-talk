package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	v1 := Point{1, 2}      // has type Vertex
	v1 = Point{X: 1, Y: 2} // Vertex{1, 2}
	v1 = Point{Y: 2, X: 1} // Vertex{1, 2}
	v1 = Point{X: 1}       // Y:0 is implicit - Vertex{1, 0}
	v1 = Point{Y: 2}       // Y:0 is implicit - Vertex{0, 2}
	v1 = Point{}           // X:0 and Y:0 - Vertex{0, 0}
	p := &Point{1, 2}      // has type *Vertex

	fmt.Println(v1.X)
	v1.X = 4
	fmt.Println(v1.X)
	fmt.Println(p.X)
}
