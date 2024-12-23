### Binary Search

Start left and right pointers at indices 0 and n-1, then find the mid point and see if that is equal to, ess than, 
or greater than your target. If nums[mid] > target, go left by moving the right pointer to mid-1. If nums[mid] < target, 
go right by moving left to mid+1. Binary Search reduces search time complexity from O(N) to O(NLogN)

When to use it?
Input is sorted and you need to find a number

Finding the position of insertion in a sorted list

Handling duplicates in sorted arrays

Searching in rotated sorted arrays

LeetCode Questions
34. [Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/)

153. [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/)

33. [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/description/)
