package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func ExampleAdd() {
	sum := Add(2, 5)
	fmt.Println(sum)
	// Output: 7
}