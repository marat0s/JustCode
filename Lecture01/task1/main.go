package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, ok := numMap[target-nums[i]]
		if ok {
			result := []int{j, i}
			return result
		}
		numMap[nums[i]] = i
	}
	answer := []int{-1, -1}
	return answer
}
