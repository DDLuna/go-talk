package main

import "fmt"

func main() {
	slice := make([]int, 0, 2)
	printSlice(slice)
	slice = append(slice, 1, 2)
	printSlice(slice)
	slice = append(slice, 3)
	printSlice(slice)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
