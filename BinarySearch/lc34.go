package main

import "fmt"

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	var result []int

	lowerBound := -1
	upperBound := -1

	if len(nums) == 1 && target == nums[0] {
		lowerBound = 0
		upperBound = 0
	}

	for left < len(nums)-1 && right > -1 {

		if nums[left] == target {
			if lowerBound == -1 {
				lowerBound = left
			}
		}
		left++

		if nums[right] == target {
			if upperBound == -1 {
				upperBound = right
			}
		}
		right--

	}

	if min(lowerBound, upperBound) == -1 {
		result = append(result, max(lowerBound, upperBound))
		result = append(result, max(lowerBound, upperBound))
	}

	if len(result) == 0 {
		result = append(result, lowerBound)
		result = append(result, upperBound)
	}

	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 3, 3, 3, 4, 5, 9}

	fmt.Println(searchRange(nums, 3))
	//nums := []int{5, 7, 7, 8, 8, 10}
	//fmt.Println(searchRange(nums, 3))

}
