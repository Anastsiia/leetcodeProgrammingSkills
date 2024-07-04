/*Given the head of a singly linked list, reverse the list, and return the reversed list.*/
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// forward solution
/*func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head.Next //2
	tempPrev := head  //1
	for curr != nil {
		tempNext := curr.Next //3
		curr.Next = tempPrev
		tempPrev = curr
		curr = tempNext
	}
	head.Next = nil
	return tempPrev
}*/

// solution with recurtion aka backwords
func reverse(head *ListNode) (h, nh *ListNode) {
	start := head
	if head.Next != nil {
		start, nh = reverse(head.Next)
		nh.Next = head
		head.Next = nil
	}
	return start, head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	rev, _ := reverse(head)
	return rev
}
