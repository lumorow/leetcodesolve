package main

import "fmt"

// Remove Duplicates from Sorted Array
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/
// Input: nums = [0,0,1,1,1,2,2,3,3,4]
// Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
// It does not matter what you leave beyond the returned k (hence they are underscores).
func removeDuplicates(nums []int) int {
	exists := make(map[int]struct{}, len(nums)/2)
	count := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		if _, ok := exists[nums[i]]; !ok {
			exists[nums[i]] = struct{}{}
			count++
		} else {
			nums = append(nums[0:i], nums[i+1:]...)
			l--
			i--
		}
	}

	return count
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
	fmt.Println(nums)
}
