/* You are given two strings s and t. String t is generated by random shuffling string s and then add one more letter at a random position. Return the letter that was added to t.*/
package main

/* cool idea to use ASCII :
func findTheDifference(s string, t string) byte {
    sumS, sumT := 0, 0
    for _, i := range s {
        sumS += int(i)
    }

    for _, i := range t {
        sumT += int(i)
    }

    diff := sumT - sumS

    return byte(diff)
} */

func findTheDifference(s string, t string) byte {
	stringMap := make(map[rune]int, len(s))
	for _, v := range s {
		stringMap[v]++
	}
	for _, v := range t {
		if stringMap[v] == 0 {
			return byte(v)
		}
		stringMap[v]--
	}
	return '0'
}
