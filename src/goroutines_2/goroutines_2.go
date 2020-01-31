package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(num int) {
			defer wg.Done()
			sayHello(num)
		}(i)
	}
	wg.Wait()
}

func sayHello(num int) {
	fmt.Println("Hello from goroutine ", num)
}
