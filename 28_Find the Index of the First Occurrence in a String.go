/* Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.*/
package main

/* we can use slices to make in more elegant:
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	h := len(haystack)
	n := len(needle)
	for i := 0; i < h-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}*/

/* also we can use func string.Index:
func strStrSample(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}*/

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	for a, v := range haystack {
		if v == rune(needle[0]) && len(haystack)-a >= len(needle) {
			i, b := 0, a
			for ; i < len(needle); i++ {
				if haystack[b] == needle[i] {
					b++
				} else {
					break
				}
			}
			if i-1 == len(needle)-1 {
				return a
			}
		}
	}
	return -1
}
