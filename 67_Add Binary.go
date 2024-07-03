/*
Given two binary strings a and b, return their sum as a binary string.
*/
package main

// solution with loop but a bit more elegant
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	res := ""
	i, j, reminder := len(a)-1, len(b)-1, 0
	for ; i >= 0; i-- {
		var bBit int
		if j < 0 {
			bBit = 0
		} else {
			bBit = int(b[j] - '0')
		}
		sum := bBit + int(a[i]-'0') + reminder
		reminder = sum / 2
		res = string(sum%2+'0') + res
		j--
	}
	if reminder == 1 {
		res = "1" + res
	}
	return res
}

//solution using string
/*func addBinary(a string, b string) string {
	if len(a) > len(b) {
		diff := len(a) - len(b)
		for i := 0; i < diff; i++ {
			b = "0" + b
		}
	} else {
		diff := len(b) - len(a)
		for i := 0; i < diff; i++ {
			a = "0" + a
		}
	}
	res := ""
	reminder := "0"
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '0' && b[i] == '0' {
			res = reminder + res
			if reminder == "1" {
				reminder = "0"
			}
		} else if a[i] == '1' && b[i] == '1' {
			res = reminder + res
			reminder = "1"
		} else {
			if reminder == "1" {
				res = "0" + res
			} else {
				res = "1" + res
				reminder = "0"
			}
		}
	}
	if reminder == "1" {
		res = "1" + res
	}

	return res
}

//solution with ParseInt & FormatInt, only for small numbers, runtime error if its bigger then int, we can use bigInt library for big ones
aInt, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		panic(err)
	}
	bInt, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		panic(err)
	}
	sum := aInt + bInt
	res := strconv.FormatInt(sum, 2)
	return res
*/
