/* Given two strings s and t, return true if t is an anagram of s, and false otherwise. An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.*/
package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[rune]int, len(s))
	for _, v := range s {
		mapS[v]++
	}
	for _, v := range t {
		mapS[v]--
		if mapS[v] < 0 {
			return false
		}
	}
	return true
}
