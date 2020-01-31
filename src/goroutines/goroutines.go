package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sayHello(i)
	}
	time.Sleep(time.Second)
}

func sayHello(number int) {
	fmt.Printf("Hello from goroutine %d\n", number)
}
