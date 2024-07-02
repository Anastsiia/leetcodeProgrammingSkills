/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.
*/
package main

func setZeroes(matrix [][]int) {
	col := -1
	for i := 0; i < len(matrix); i++ {
		for q := 0; q < len(matrix[0]); q++ {
			if matrix[i][q] == 0 {
				if i == 0 {
					matrix[0][0] = 0
				}
				if q == 0 {
					col = 0
				} else {
					matrix[0][q] = 0
					matrix[i][0] = 0
				}
			}
		}
	}
	for a := 1; a < len(matrix[0]); a++ {
		if matrix[0][a] == 0 {
			for c := 1; c < len(matrix); c++ {
				matrix[c][a] = 0
			}
		}
	}
	for b := 1; b < len(matrix); b++ {
		if matrix[b][0] == 0 {
			for c := 1; c < len(matrix[0]); c++ {
				matrix[b][c] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for c := 0; c < len(matrix[0]); c++ {
			matrix[0][c] = 0
		}
	}
	if col == 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

// 111
// 101
// 111

// 101
// 000
// 101
