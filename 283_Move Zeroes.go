/* Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements. Note that you must do this in-place without making a copy of the array.*/
package main

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for a := i + 1; a < len(nums); a++ {
				if nums[a] != 0 {
					nums[i], nums[a] = nums[a], nums[i]
					break
				}
			}
		}
	}
}
