package main

import "fmt"

func subsets(nums []int) [][]int {
	res, sol := [][]int{}, []int{}

	backtrack(0, &res, sol, nums)
	return res
}

func backtrack(i int, res *[][]int, sol []int, nums []int) {
	if i == len(nums) {
		temp := make([]int, len(sol))
		copy(temp, sol)
		(*res) = append(*res, temp)
		return
	}

	backtrack(i+1, res, sol, nums)

	sol = append(sol, nums[i])
	backtrack(i+1, res, sol, nums)
	sol = sol[:len(sol)-1]
}

func main() {
	nums := []int{1, 2, 3}

	fmt.Println(subsets(nums))
	//nums := []int{5, 7, 7, 8, 8, 10}
	//fmt.Println(searchRange(nums, 3))

}
