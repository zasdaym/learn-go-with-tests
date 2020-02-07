package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Robot", "")
		want := "Hello, Robot"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying in Spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, world"
		assertCorrectMessage(t, got, want)
	})
}