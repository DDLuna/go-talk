package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for {
		fmt.Println("Infinite loop!!!")
		if time.Now().Sub(start).Milliseconds() > 200 {
			break
		}
	}
}