package main

import (
	"fmt"
	"strings"
)

func characterReplacement(s string, replacementLimit int) int {
	maxCount := 0
	letters := strings.Split(s, "")
	letterCounter := make(map[string]int)
	left, right, result := 0, 0, 0
	for right = 0; right < len(s); right++ {
		letterCounter[letters[right]]++
		//set maxCount which sums the dupe letter count in the window
		maxCount = max(maxCount, letterCounter[letters[right]])

		//if window size minus maxCount of dupe is more than replacement Limit, advance left pointer
		if right-left+1-maxCount > replacementLimit {
			letterCounter[letters[left]]--
			left++
		}
		//result is the max window size
		result = max(result, right-left+1)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))

}
