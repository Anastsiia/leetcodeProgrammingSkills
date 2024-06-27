/*
An array is monotonic if it is either monotone increasing or monotone decreasing.

An array nums is monotone increasing if for all i <= j, nums[i] <= nums[j]. An array nums is monotone decreasing if for all i <= j, nums[i] >= nums[j].

Given an integer array nums, return true if the given array is monotonic, or false otherwise.
*/
package main

func isMonotonic(nums []int) bool {
	isIncreasing := true
	isDecreasing := true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > 0 {
			isDecreasing = false
		} else if nums[i+1]-nums[i] < 0 {
			isIncreasing = false
		}
		if !isIncreasing && !isDecreasing {
			return false
		}
	}

	return isDecreasing || isIncreasing
}

// func isMonotonic(nums []int) bool {
// 	len := len(nums)
// 	if len <= 2 {
// 		return true
// 	}
// 	if nums[0] <= nums[len-1] {
// 		for i := 1; i < len; i++ {
// 			if nums[i] < nums[i-1] {
// 				return false
// 			}
// 		}
// 	} else {
// 		for i := 1; i < len; i++ {
// 			if nums[i] > nums[i-1] {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
