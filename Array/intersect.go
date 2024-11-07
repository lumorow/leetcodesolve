package main

// Intersection of Two Arrays II
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/674/

// Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]
// Example 2:

// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]
// Explanation: [9,4] is also accepted.

// Follow up:
// What if the given array is already sorted? How would you optimize your algorithm?
// What if nums1's size is small compared to nums2's size? Which algorithm is better?
// What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	res := make([]int, 0, len(nums1))

	for _, num := range nums1 {
		m[num]++
	}

	for _, num := range nums2 {
		if _, ok := m[num]; ok {
			count := m[num]
			if count > 0 {
				m[num]--
				res = append(res, num)
			}
		}
	}

	return res
}
