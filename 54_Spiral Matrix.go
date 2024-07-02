/*
Given an m x n matrix, return all elements of the matrix in spiral order.
*/
package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var res []int
	wight := len(matrix[0])
	hight := len(matrix)

	for i := 0; len(res) < hight*wight; i++ {
		for r := i; r < wight-i && len(res) < hight*wight; r++ {
			fmt.Println(1)
			res = append(res, matrix[i][r])
		}
		for d := i + 1; d < hight-i && len(res) < hight*wight; d++ {
			fmt.Println(2)
			res = append(res, matrix[d][wight-i-1])
		}
		for l := wight - i - 2; l >= i && len(res) < hight*wight; l-- {
			res = append(res, matrix[hight-i-1][l])
		}
		for u := hight - i - 2; u > i && len(res) < hight*wight; u-- {
			fmt.Println(4)
			res = append(res, matrix[u][i])
		}
	}
	return res
}

//[1,2,3,6,9,8,7,4,5]
//[1,2,3,4,8,12,11,10,9,5,6,7]
