package main

import (
	"testing"
)

func TestCompareSlices(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3, 2, 6}
	expected := false

	result := compareSlices(slice1, slice2)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
