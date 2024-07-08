/*
You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

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
//stack solution
type listStack struct {
	items []int
	len   int
}

func (s *listStack) isEmpty() bool {
	return s.len == 0
}

func (s *listStack) Push(elment int) {
	s.items = append(s.items, elment)
	s.len++
}

func (s *listStack) Pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	res := 0
	index := s.len - 1
	res = s.items[index]
	s.items = s.items[:index]
	s.len--
	return res, true
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = &ListNode{5, nil}
	l2 = &ListNode{5, nil}
	s1, s2, s3 := listStack{}, listStack{}, listStack{}
	for ; l1 != nil; l1 = l1.Next {
		s1.Push(l1.Val)
	}
	for ; l2 != nil; l2 = l2.Next {
		s2.Push(l2.Val)
	}
	reminder := 0
	for !s1.isEmpty() || !s2.isEmpty() {
		num1, success := s1.Pop()
		if success {
			reminder += num1
		}
		num2, success := s2.Pop()
		if success {
			reminder += num2
		}
		s3.Push(reminder % 10)
		reminder /= 10
	}
	if reminder != 0 {
		s3.Push(reminder)
	}
	res := new(ListNode)
	for list := res; !s3.isEmpty(); list = list.Next {
		num, _ := s3.Pop()
		list.Next = &ListNode{num, nil}
	}
	return res.Next
}

/* using my previous functions
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
	return reverseList(addTwoNumbers(l1, l2))
}*/
