package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// // 28:
	// word1 := "a"
	// word2 := "a"
	// fmt.Println(strStr(word1, word2))

	//66:
	// nums := []int{9, 9}
	// fmt.Println(plusOne(nums))

	//242
	// word1 := "anagram"
	// word2 := "zagaram"
	// fmt.Println(isAnagram(word1, word2))

	//283:
	// nums := []int{0, 1, 0, 3, 12}
	// moveZeroes(nums)
	// fmt.Println(nums)

	//389:
	// word1 := "a"
	// word2 := "aa"
	// fmt.Println(findTheDifference(word1, word2))

	//458:
	// word3 := "ababab"
	// fmt.Println(repeatedSubstringPattern(word3))

	//1502:
	// nums := []int{7, 5, 1}
	// fmt.Println(canMakeArithmeticProgression(nums))

	//1768:
	// word1 := "a b c"
	// word2 := "xyz"
	// fmt.Println(mergeAlternately(word1, word2))

	//1822:
	// nums := []int{-1, 0, 1}
	// fmt.Println(arraySign(nums))

	//896:
	// nums := []int{-1, 0, 0}
	// fmt.Println(isMonotonic(nums))

	//13
	// str := "III"
	// fmt.Println(romanToInt(str))

	//58
	// str := "   fly me   to   the moon  "
	// fmt.Println(lengthOfLastWord(str))

	//709:
	// str := "III"
	// fmt.Println(toLowerCase(str))

	//682:
	// ops := []string{"1", "C"}
	// fmt.Println(calPoints(ops))

	//657:
	// moves := "LL"
	// fmt.Println(judgeCircle(moves))

	//1275:
	// moves := [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}
	// fmt.Println(tictactoe(moves))

	//1041:
	// instructions := "GL"
	// fmt.Println(isRobotBounded(instructions))

	//1672:
	// accounts := [][]int{{1, 12, 3}, {3, 2, 1}}
	// fmt.Println(maximumWealth(accounts))

	//1572
	// mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// fmt.Println(diagonalSum(mat))

	//54:
	// mat := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	// fmt.Println(spiralOrder(mat))

	//73:
	// mat := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	// setZeroes(mat)
	// fmt.Println(mat)

	//1523:
	// low := 3
	// high := 9
	// fmt.Println(countOdds(low, high))

	//1491:
	// nums := []int{7, 5, 1}
	// fmt.Println(average(nums))

	//860:
	// nums := []int{5, 5, 5, 5, 10, 5, 10, 10, 10, 20}
	// fmt.Println(lemonadeChange(nums))

	//976:
	// nums := []int{1, 2, 1, 10}
	// fmt.Println(largestPerimeter(nums))

	//1232:
	// mat := [][]int{{1, 2}, {2, 3}, {3, 5}}
	// fmt.Println(checkStraightLine(mat))

	//67:
	// word1 := "11"
	// word2 := "1"
	// fmt.Println(addBinary(word1, word2))

	//43:
	// word1 := "140"
	// word2 := "721"
	// fmt.Println(multiply(word1, word2))

	//50:
	// x := 2.0
	// n := 4
	// fmt.Println(myPow(x, n))

	//21:
	// node3 := ListNode{4, nil}
	// node2 := ListNode{2, &node3}
	// node1 := ListNode{1, &node2}
	// node6 := ListNode{4, nil}
	// node5 := ListNode{3, &node6}
	// node4 := ListNode{1, &node5}
	// list1 := &node1
	// list2 := &node4
	// // for list1 != nil {
	// // 	fmt.Println(list1.Val)
	// // 	list1 = list1.Next
	// // }
	// // for list2 != nil {
	// // 	fmt.Println(list2.Val)
	// // 	list2 = list2.Next
	// // }
	// res := mergeTwoLists(list1, list2)
	// for res != nil {
	// 	fmt.Println(res.Val)
	// 	res = res.Next
	// }

	//206:
	// node3 := ListNode{4, nil}
	// node2 := ListNode{2, &node3}
	// node1 := ListNode{1, &node2}
	// list1 := &node1
	// res := reverseList(list1)
	// for res != nil {
	// 	fmt.Println(res.Val)
	// 	res = res.Next
	// }

	//2:
	// node3 := ListNode{3, nil}
	// node2 := ListNode{4, nil}
	// node1 := ListNode{2, &node2}
	// node6 := ListNode{4, nil}
	// node5 := ListNode{6, &node6}
	// node4 := ListNode{5, &node5}
	// list1 := &node1
	// list2 := &node4
	// // for list1 != nil {
	// // 	fmt.Println(list1.Val)
	// // 	list1 = list1.Next
	// // }
	// // for list2 != nil {
	// // 	fmt.Println(list2.Val)
	// // 	list2 = list2.Next
	// // }
	// res := addTwoNumbers(list1, list2)
	// for res != nil {
	// 	fmt.Println(res.Val)
	// 	res = res.Next
	// }

	//445:
	// node3 := ListNode{3, nil}
	// node2 := ListNode{4, nil}
	// node1 := ListNode{2, &node2}
	// node6 := ListNode{4, nil}
	// node5 := ListNode{6, &node6}
	// node4 := ListNode{5, &node5}
	// list1 := &node1
	// list2 := &node4
	// // for list1 != nil {
	// // 	fmt.Println(list1.Val)
	// // 	list1 = list1.Next
	// // }
	// // for list2 != nil {
	// // 	fmt.Println(list2.Val)
	// // 	list2 = list2.Next
	// // }
	// res := addTwoNumbers2(list1, list2)
	// for res != nil {
	// 	fmt.Println(res.Val)
	// 	res = res.Next
	// }
}
