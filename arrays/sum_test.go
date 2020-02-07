package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{1, 2})
		want := []int{9, 2}
		checkSums(t, got, want)
	})

	t.Run("safely sums empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2})
		want := []int{0, 2}
		checkSums(t, got, want)
	})

}
