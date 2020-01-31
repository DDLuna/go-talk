package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	num, err := strconv.Atoi("12345")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(num)
}

func addPositiveNumbers(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New("Both numbers must be positive")
	}
	return a + b, nil
}
