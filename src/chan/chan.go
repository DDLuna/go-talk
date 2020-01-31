package main

import "fmt"

func main() {
	ch := make(chan string)
	routines := 3

	for i := 0; i < routines; i++ {
		go func(num int) {
			sayHello(num, ch)
		}(i)
	}
	for i := 0; i < routines; i++ {
		<-ch
	}
}

func sayHello(num int, ch chan string) {
	fmt.Println("Hello goroutine ", num)
	ch <- "done"
}
