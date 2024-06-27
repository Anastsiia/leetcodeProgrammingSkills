/*
	There is a function signFunc(x) that returns:

1 if x is positive.
-1 if x is negative.
0 if x is equal to 0.
You are given an integer array nums. Let product be the product of all values in the array nums.

Return signFunc(product).
*/
package main

func arraySign(nums []int) int {
	res := 1
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		if v < 0 {
			res *= -1
		}
	}
	return res
}
