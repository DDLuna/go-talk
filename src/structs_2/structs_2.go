package main

import "fmt"

type tree struct {
	value int
	left  *tree
	right *tree
}

func (t *tree) add(newValue int) {
	if t == nil {
		t = &tree{value: newValue}
	}
	if t.value > newValue {
		t.right.add(newValue)
	} else {
		t.left.add(newValue)
	}
}

func (t *tree) printAsc() {
	if t == nil {
		return
	}
	t.left.printAsc()
	fmt.Println(t.value)
	t.right.printAsc()
}

func main() {

}
