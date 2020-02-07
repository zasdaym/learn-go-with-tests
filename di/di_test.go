package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Robot")

	got := buffer.String()
	want := "Hello, Robot"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}