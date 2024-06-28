/*
Given a string s, return the string after replacing every uppercase letter with the same lowercase letter.
*/
package main

// import "strings"

func toLowerCase(s string) string {
	/*we can use build-in func:
	return strings.ToLower(s)*/

	str := []rune(s)
	for i, v := range str {
		if v >= 'A' && v <= 'Z' {
			str[i] = v + 32
		}
	}
	return string(str)
}
