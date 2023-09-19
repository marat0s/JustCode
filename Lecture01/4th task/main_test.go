package main

import (
	"reflect"
	"testing"
)

func TestIntSliceSort(t *testing.T) {
	s := intSlice{slice: []int{4, 3, 2, 5, 1}}
	expected := []int{1, 2, 3, 4, 5}

	s.sort()

	if !reflect.DeepEqual(s.slice, expected) {
		t.Errorf("Expected %v, but got %v", expected, s.slice)
	}
}
