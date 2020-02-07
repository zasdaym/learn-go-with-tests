package main

import (
	"testing"
)

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("Should be find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error")
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "this is test"
		newDefinition := "this is new"
		dictionary := Dictionary{word: definition}
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Non existing word", func(t *testing.T) {
		word := "test"
		definition := "this is test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Update("unknown", "unknown")

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "this is test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Delete(word)

		assertError(t, err, nil)
	})

	t.Run("Non existing word", func(t *testing.T) {
		word := "test"
		definition := "this is test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Delete("unknown")

		assertError(t, err, ErrWordDoesNotExist)
	})
}
