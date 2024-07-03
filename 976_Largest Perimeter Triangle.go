/*
Given an integer array nums, return the largest perimeter of a triangle with a non-zero area, formed from three of these lengths. If it is impossible to form any triangle of a non-zero area, return 0.
*/
package main

import (
	"sort"
)

func largestPerimeter(nums []int) int {
	len := len(nums)
	sort.Ints(nums)
	j := 0
	for i := len - 1; i >= 2 && j < 45; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
		j++
	}
	return 0
}

//1,2,1,10
