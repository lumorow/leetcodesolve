package main

// Rotate Array
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

// Example 1:
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]

// Example 2:
// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]

func rotate(nums []int, k int) {
	if k > 0 && len(nums) > 1 {

		if k >= len(nums) {
			k = k - (len(nums) * (k / len(nums)))
		}

		nums2 := append(nums[len(nums)-k:], nums[0:len(nums)-k]...)

		for i := range nums {
			nums[i] = nums2[i]
		}
	}
}
