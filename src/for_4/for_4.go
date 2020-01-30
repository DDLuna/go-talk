package main

import (
	"fmt"
)

func main() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	for index, value := range arr {
		fmt.Printf("index: %d, value: %v\n", index, value)
	}
}
