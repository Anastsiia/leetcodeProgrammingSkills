/*
	You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	curr := &ListNode{0, nil}
	res := curr
	reminder := 0
	for l1 != nil || l2 != nil {
		num1, num2 := 0, 0
		if l1 == nil {
			num1 = 0
		} else if l1 != nil {
			num1 = l1.Val
		}
		if l2 == nil {
			num2 = 0
		} else if l2 != nil {
			num2 = l2.Val
		}
		curr.Val = num1 + num2 + reminder
		if curr.Val > 9 {
			reminder = curr.Val / 10
			curr.Val = curr.Val % 10
		} else {
			reminder = 0
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if reminder == 0 && (l1 == nil && l2 == nil) {
			return res
		}
		temp := &ListNode{0, nil}
		curr.Next = temp
		curr = temp

	}
	if reminder != 0 {
		curr.Val = reminder
	}
	return res
}
