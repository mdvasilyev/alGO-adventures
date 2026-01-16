package main

import (
	"fmt"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil && list2 != nil {
		cur.Next = &ListNode{}
		cur = cur.Next
		if list1.Val <= list2.Val {
			cur.Val = list1.Val
			list1 = list1.Next
		} else {
			cur.Val = list2.Val
			list2 = list2.Next
		}
	}
	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}
	return head.Next
}

func main() {
	l1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l2 := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}

	r := mergeTwoLists(&l1, &l2)
	fmt.Println(r)
}
