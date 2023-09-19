package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	result := twoSum(nums, target)
	expected := []int{1, 2}

	if !equalSlices(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
