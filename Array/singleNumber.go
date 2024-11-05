package main

// Single Number
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/
//
// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space.

//Example 1:
//Input: nums = [2,2,1]
//Output: 1

//Example 2:
//Input: nums = [4,1,2,1,2]
//Output: 4

//Example 3:
//Input: nums = [1]
//Output: 1

func singleNumber(nums []int) int {
	res := nums[0]
	m := make(map[int]int, len(nums))

	for _, num := range nums {
		m[num]++
	}

	for key, value := range m {
		if value == 1 {
			res = key
		}
	}

	return res
}
