package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	// Create a letterCounter dictionary
	letters := strings.Split(s, "")
	letterCounter := make(map[string]int)
	left, right, result := 0, 0, 0

	// Slide the window to the right
	for right < len(letters) {
		r := (letters[right]) //get the letter
		letterCounter[(r)]++  //count it

		right++ //step right
		// If there are duplicates between left and right, advance the left pointer and decrement the letter count
		for letterCounter[(r)] > 1 { //if the counter is more than one it's a dupe
			l := (letters[left]) //get the left pointer letter
			letterCounter[(l)]-- //decrement the count
			left++               //move left pointer to the right
		}
		// Update the longest substring length
		if result < right-left {
			result = right - left
		}
	}
	return result

}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))

}
