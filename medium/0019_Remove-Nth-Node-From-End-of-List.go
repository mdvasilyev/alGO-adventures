package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Time:	O(т)
// Space:	O(n)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{Val: 0, Next: head}
	cur := res
	for i := 0; i <= n; i++ {
		cur = cur.Next
	}
	lag := res
	for cur != nil {
		cur = cur.Next
		lag = lag.Next
	}
	lag.Next = lag.Next.Next
	return res.Next
}

func main() {
	head1, n1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}, 2
	head2, n2 := &ListNode{Val: 1, Next: nil}, 1
	head3, n3 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}, 1

	fmt.Println(removeNthFromEnd(head1, n1))
	fmt.Println(removeNthFromEnd(head2, n2))
	fmt.Println(removeNthFromEnd(head3, n3))
}
