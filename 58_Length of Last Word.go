/*
Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal
substring

	consisting of non-space characters only.
*/
package main

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	return len(words[len(words)-1])
}

/* solution without build-in functions
func lengthOfLastWord(s string) int {
	i, res := len(s)-1, 0
	for ; i >= 0; i-- {
		if s[i] == ' ' && res == 0 {
			continue
		} else {
			if s[i] != ' ' {
				res++
			} else {
				break
			}
		}
	}
	return res
}
*/
