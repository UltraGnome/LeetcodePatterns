package main

import (
	"fmt"
	"strings"
)

func countGoodSubstrings(s string) int {
	letters := strings.Split(s, "")
	left, right, result := 0, 2, 0

	for right < len(s) {
		//get substring
		subby := letters[left] + letters[right-1] + letters[right]
		//is all unique?  count it
		if !hasDuplicateLetters(subby) {
			result++
		}

		left++
		right++

	}
	return result
}

func hasDuplicateLetters(input_value string) bool {
	result := false
	counts := make(map[rune]int)
	for _, char := range input_value {
		if counts[char] == 1 {
			result = true
			return result
		}
		counts[char]++ //increment the count
	}
	return result
}

func main() {
	fmt.Println(countGoodSubstrings("xyzzaz"))

}
