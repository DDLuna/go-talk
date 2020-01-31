package main

import "testing"

func TestArea(t *testing.T) {
	rect := Rectangle{4.0, 3.0}
	want := 12.0
	got := rect.Area()
	if want != got {
		t.Errorf("got %f want %f", got, want)
	}
}
