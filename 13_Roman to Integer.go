/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
*/
package main

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := len(s) - 1; i >= 0; i-- {
		if i == 0 {
			res += dict[s[i]]
		} else if dict[s[i]] <= dict[s[i-1]] {
			res += dict[s[i]]
		} else {
			res += dict[s[i]] - dict[s[i-1]]
			i--
		}
	}
	return res
}

/* standart solution with conditions
func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'M':
			res += 1000
		case 'D':
			res += 500
		case 'C':
			if i != len(s)-1 && s[i+1] == 'D' {
				res += 400
				i++
			} else if i != len(s)-1 && s[i+1] == 'M' {
				res += 900
				i++
			} else {
				res += 100
			}
		case 'L':
			res += 50
		case 'X':
			if i != len(s)-1 && s[i+1] == 'L' {
				res += 40
				i++
			} else if i != len(s)-1 && s[i+1] == 'C' {
				res += 90
				i++
			} else {
				res += 10
			}
		case 'V':
			res += 5
		case 'I':
			if i != len(s)-1 {
				if s[i+1] == 'V' {
					res += 4
					i++
					break
				} else if s[i+1] == 'X' {
					res += 9
					i++
					break
				}
			}
			res += 1
		}
	}
	return res
}
*/
