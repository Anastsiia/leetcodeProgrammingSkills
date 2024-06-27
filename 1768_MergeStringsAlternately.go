/* You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string. Return the merged string.*/
package main

func mergeAlternately(word1 string, word2 string) string {
	var res string
	maxLen := len(word1)
	if len(word1) < len(word2) {
		maxLen = len(word2)
	}
	for i := 0; i < maxLen; i++ {
		if i < len(word1) {
			res = res + string(word1[i])
		}
		if i < len(word2) {
			res = res + string(word2[i])
		}
	}
	return res
}
