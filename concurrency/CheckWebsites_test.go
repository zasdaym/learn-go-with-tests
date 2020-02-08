package main

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://mock.mock" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://example.com",
		"http://mock.mock",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://example.com": true,
		"http://mock.mock": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}