/* Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.*/
package main

import (
	"fmt"
)

func repeatedSubstringPattern(s string) bool {
	// lenStr := len(s)

	// if lenStr == 1 {
	// 	return false
	// }
	// for i := 2; i <= lenStr; i++ {
	// 	if lenStr%i == 0 {
	// 		for a := 0; a < lenStr-1; {
	// 			if s[a:a+lenStr/i] != s[a+lenStr/i:a+2*lenStr/i] {
	// 				break
	// 			}
	// 			a = a + lenStr/i
	// 			if a+lenStr/i > lenStr-1 {
	// 				return true
	// 			}
	// 		}
	// 	}
	// }
	// return false

	ss := s
	fmt.Println(ss)

	for i := len(ss) - 1; i > 0; i-- {
		fmt.Println("i %v", i)
		s = string(s[len(s)-1]) + s[0:len(ss)-1]
		fmt.Println("s %v", s)
		if s == ss {
			return true
		}
	}
	return false

}

func main() {
	word1 := "ababab"

	fmt.Println(repeatedSubstringPattern(word1))
}
