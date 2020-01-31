package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	f, err := os.Create("defer.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	/*

	 */
}
