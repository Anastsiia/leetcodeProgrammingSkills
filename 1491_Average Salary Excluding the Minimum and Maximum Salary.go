/*
You are given an array of unique integers salary where salary[i] is the salary of the ith employee.

Return the average salary of employees excluding the minimum and maximum salary. Answers within 10-5 of the actual answer will be accepted.
*/
package main

import "golang.org/x/exp/slices"

func average(salary []int) float64 {
	res := 0.0
	slices.Sort(salary)
	//also can use
	// sort.Ints(salary)
	for i := 1; i < len(salary)-1; i++ {
		res += float64(salary[i])
	}
	res = res / float64((len(salary) - 2))
	return res
}
