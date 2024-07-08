/*
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/
package main

import (
	"strconv"
)

func multiplyByOneDigit(num string, digit byte, saveTo []string, index int) {
	res := ""
	digitInt := int(digit - 48)
	reminder := 0
	for i := len(num) - 1; i >= 0; i-- {
		curr := int(num[i]-48)*digitInt + reminder
		if curr > 9 {
			reminder = curr / 10
			curr = curr % 10
		} else {
			reminder = 0
		}
		res = strconv.Itoa(curr) + res
	}
	if reminder != 0 {
		res = strconv.Itoa(reminder) + res
	}
	for i := index; i > 0; i-- {
		res = res + "0"
	}
	saveTo[index] = res
}

func sumArr(arr []string) string {
	res := ""
	goalLen := len(arr[len(arr)-1])
	for i := 0; i < len(arr); i++ {
		for a := goalLen - len(arr[i]); a > 0; a-- {
			arr[i] = "0" + arr[i]
		}
	}
	reminder := 0
	for i := goalLen - 1; i >= 0; i-- {
		num := 0
		for _, v := range arr {
			num += int(v[i] - 48)
		}
		num += reminder
		res = strconv.Itoa(num%10) + res
		reminder = num / 10
	}
	if reminder != 0 {
		res = strconv.Itoa(reminder) + res
	}
	return res
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	elem := make([]string, len(num2))
	j := 0
	for i := len(num2) - 1; i >= 0; i-- {
		multiplyByOneDigit(num1, num2[i], elem, j)
		j++
	}
	res := sumArr(elem)
	return res

}

// solution with my Atoi & Itoa. does not work with big numbers!
/*func myAtoi(str string) uint64 {
	var res uint64
	for i := 0; i < len(str); i++ {
		res = res*10 + uint64(str[i]-48)
	}
	fmt.Println(res)
	return res
}

func myItoa(num uint64) string {
	res := ""
	if num == 0 {
		return "0"
	}
	for num > 0 {
		res = strconv.FormatUint((num%10), 10) + res
		num /= 10
	}
	return res
}

func multiply(num1 string, num2 string) string {
	return myItoa(myAtoi(num1) * myAtoi(num2))
}
*/
