/*
You are given an array of unique integers salary where salary[i] is the salary of the ith employee.

Return the average salary of employees excluding the minimum and maximum salary. Answers within 10-5 of the actual answer will be accepted.
*/
package main

import "golang.org/x/exp/slices"

func average(salary []int) float64 {
	res := 0.0

	// slices.Sort - is better to use when you need to sort different type of slice(strings, float)
	slices.Sort(salary)

	//also can use. is better because From the standard sort package in the core library.is preferred when you specifically need to sort integer
	// sort.Ints(salary)
	for i := 1; i < len(salary)-1; i++ {
		res += float64(salary[i])
	}
	res = res / float64((len(salary) - 2))
	return res
}
