/*
You are keeping the scores for a baseball game with strange rules. At the beginning of the game, you start with an empty record.

You are given a list of strings operations, where operations[i] is the ith operation you must apply to the record and is one of the following:

An integer x.
Record a new score of x.
'+'.
Record a new score that is the sum of the previous two scores.
'D'.
Record a new score that is the double of the previous score.
'C'.
Invalidate the previous score, removing it from the record.
Return the sum of all the scores on the record after applying all the operations.

The test cases are generated such that the answer and all intermediate calculations fit in a 32-bit integer and that all operations are valid.
*/
package main

import (
	"strconv"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Double() {
	s.items = append(s.items, s.items[len(s.items)-1]*2)
}

func (s *Stack) Add() {
	s.items = append(s.items, s.items[len(s.items)-1]+s.items[len(s.items)-2])
}

func (s *Stack) Pop() {
	s.items = append(s.items[:len(s.items)-1])
}

func (s *Stack) Sum() int {
	res := 0
	for _, v := range s.items {
		res += v
	}
	return res
}

func calPoints(operations []string) int {
	// solution with a stack
	var records Stack
	for _, v := range operations {
		switch v {
		case "+":
			records.Add()
		case "C":
			records.Pop()
		case "D":
			records.Double()
		default:
			data, _ := strconv.Atoi(v)
			records.Push(data)
		}
	}
	return records.Sum()
	/* solution with an array
	records := []int{}
	res := 0
	for _, v := range operations {
		switch v {
		case "C":
			records = append(records[:len(records)-1])
		case "D":
			records = append(records, records[len(records)-1]*2)
		case "+":
			records = append(records, records[len(records)-1]+records[len(records)-2])
		default:
			curr, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			records = append(records, curr)
		}
	}
	for _, v := range records {
		res += v
	}
	return res*/
}
