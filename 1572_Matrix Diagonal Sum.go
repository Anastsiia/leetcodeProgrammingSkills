/*
Given a square matrix mat, return the sum of the matrix diagonals.

Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.
*/
package main

func diagonalSum(mat [][]int) int {
	res := 0
	matLen := len(mat)
	for i, v := range mat {
		res += v[i] + v[matLen-1-i]
	}
	if matLen%2 != 0 {
		middle := mat[matLen/2][matLen/2]
		res -= middle
	}
	return res
}

// [[1,2,3],
//  [4,5,6],
//  [7,8,9]]
