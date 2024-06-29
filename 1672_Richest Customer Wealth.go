/*
You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.

A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.
*/
package main

import "sync"

func maximumWealth(accounts [][]int) int {
	max := 0
	sums := make(chan int, len(accounts))
	wg := sync.WaitGroup{}

	for _, v := range accounts {
		wg.Add(1)
		go func(nums []int) {
			res := 0
			for _, v := range nums {
				res += v
			}
			sums <- res
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(sums)
	for v := range sums {
		if v > max {
			max = v
		}
	}
	return max
}
