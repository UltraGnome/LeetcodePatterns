package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxWater := 0

	p1, p2 := 0, len(height)-1

	for p1 < p2 {
		var minHeight int
		if height[p1] > height[p2] {
			minHeight = height[p2]
			p2--
		} else {
			minHeight = height[p1]
			p1++
		}

		thisWater := minHeight * (p2 - p1 + 1)

		if thisWater > maxWater {
			maxWater = thisWater
		}

	}

	return maxWater

}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) //49
	fmt.Println(maxArea([]int{1, 1}))                      //1
	fmt.Println(maxArea([]int{1, 2, 1}))                   //2
	fmt.Println(maxArea([]int{1, 2, 4, 3}))                //4
	fmt.Println(maxArea([]int{8, 7, 2, 1}))                //4

}

//You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
//Find two lines that together with the x-axis form a container, such that the container contains the most water.
//Return the maximum amount of water a container can store.
//Notice that you may not slant the container.
//Example 1:
//Input: height = [1,8,6,2,5,4,8,3,7]
//Output: 49
//Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
//Example 2:
//
//Input: height = [1,1]
//Output: 1
