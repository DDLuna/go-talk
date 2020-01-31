package main

import "fmt"

func main() {
	defer logPanic()
	var a *int
	// null pointer
	fmt.Println(*a)
}

func logPanic() {
	if ok := recover(); ok != nil {
		fmt.Println("null prt is scary! PANIC ", ok)
		panic(ok)
	}
}
