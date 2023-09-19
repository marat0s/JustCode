package main

import "fmt"

func compareSlices(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	sliceMap1 := make(map[int]int)
	sliceMap2 := make(map[int]int)

	for _, j := range slice1 {
		sliceMap1[j]++
	}

	for _, j := range slice2 {
		sliceMap2[j]++
	}

	for i, j := range sliceMap1 {
		if sliceMap2[i] != j {
			return false
		}
	}

	return true
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3, 2, 6}

	fmt.Println(compareSlices(slice1, slice2))
}
