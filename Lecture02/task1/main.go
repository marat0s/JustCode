package main

import "fmt"

func main() {
	fmt.Println(intToRoman(1337))
}

func intToRoman(num int) string {
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	finalResult := ""

	for i := 0; i < len(value); i++ {
		for num >= value[i] {
			num -= value[i]
			finalResult += symbols[i]
		}
	}

	return finalResult
}
