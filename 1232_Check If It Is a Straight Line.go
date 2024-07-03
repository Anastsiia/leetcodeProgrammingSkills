/*
You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point. Check if these points make a straight line in the XY plane.
*/
package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	for i := 0; i < len(coordinates)-2; i++ {
		diff1 := (coordinates[i+1][1] - coordinates[i][1]) * (coordinates[i+2][0] - coordinates[i+1][0])
		diff2 := (coordinates[i+2][1] - coordinates[i+1][1]) * (coordinates[i+1][0] - coordinates[i][0])
		fmt.Println(diff1, diff2)
		if diff1 != diff2 {
			return false
		}
	}
	return true
}

// y2 - y1   y3 - y2
// ------- = -------
// x2 - x1   x3 - x2

// y2-y1 * x3-x2 = y3-y2 * x2-x1
