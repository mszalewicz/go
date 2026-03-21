package main

// Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.
//
// Example 1:
//
// Input: nums = [1, 2, 3, 3]
//
// Output: true
//
//
// Example 2:
//
// Input: nums = [1, 2, 3, 4]
//
// Output: false

func hasDuplicate(nums []int) bool {
	seen := make(map[int]int)

	for _, number := range nums {
		if _, ok := seen[number]; ok {
			return true
		}
		seen[number] = 1
	}

	return false
}
