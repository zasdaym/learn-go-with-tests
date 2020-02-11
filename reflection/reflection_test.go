package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	} {
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Robot"},
			[]string{"Robot"},
		},
		{
			"Struct with two string field",
			struct {
				Name string
				City string
			}{"Robot", "Olympus"},
			[]string{"Robot", "Olympus"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age int
			}{"Robot", 20},
			[]string{"Robot"},
		},
		{
			"Nested fields",
			Person{
				"Robot",
				Profile{20, "Olympus"},
			},
			[]string{"Robot", "Olympus"},
		},
		{
			"Pointers",
			&Person{
				"Robot",
				Profile{20, "Olympus"},
			},
			[]string{"Robot", "Olympus"},
		},
		{
			"Slices",
			[]Profile {
				{20, "Olympus"},
				{21, "Alexandria"},
			},
			[]string{"Olympus", "Alexandria"},
		},
		{
			"Arrays",
			[2]Profile {
				{20, "Olympus"},
				{21, "Alexandria"},
			},
			[]string{"Olympus", "Alexandria"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]string{
			"foo":"bar",
			"baz":"boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}