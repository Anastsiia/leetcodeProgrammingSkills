/*
Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).
*/
package main

//O(n)-O(1) solution
/*func countOdds(low int, high int) int {
	res := 0
	for i := low; i <= high; i++ {
		//if i%2 != 0 {
		if i&1 != 0 {
			res++
		}
	}
	return res
}*/

// O(1) - O(1) bitwise solution
func countOdds(low int, high int) int {
	return ((high - low) / 2) + (low%2 | high%2)
}
