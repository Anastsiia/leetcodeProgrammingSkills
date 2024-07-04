/*
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
*/
package main

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	res := 1.0
	if n < 0 {
		return 1 / myPow(x, -n)
	}

	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n = n >> 1
	}
	return res
}
