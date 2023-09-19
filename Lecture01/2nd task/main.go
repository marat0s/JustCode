package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	prefix := str[0]
	for _, str := range str {
		for !strings.HasPrefix(str, prefix) { //HasPrefix to find common prefixes
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

func main() {
	str := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(str))
}
