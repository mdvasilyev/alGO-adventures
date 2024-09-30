package main

import (
	"fmt"
)

// ListNode Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// Time:	O(len(head))
// Space:	O(len(head))
func reverseList(head *ListNode) *ListNode {
	values := make([]int, 0, 10)
	cur := head
	for cur != nil {
		values = append(values, cur.Val)
		cur = cur.Next
	}
	res := &ListNode{}
	cur = res
	for i := len(values) - 1; i >= 0; i-- {
		cur.Next = &ListNode{}
		cur = cur.Next
		cur.Val = values[i]
	}
	return res.Next
}

func main() {
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	head2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	head3 := &ListNode{Val: 0, Next: nil}

	fmt.Println(reverseList(head1))
	fmt.Println(reverseList(head2))
	fmt.Println(reverseList(head3))
}
