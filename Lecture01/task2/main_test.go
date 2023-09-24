package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	str := []string{"flower", "flour", "flow"}
	expected := "fl"

	result := longestCommonPrefix(str)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
