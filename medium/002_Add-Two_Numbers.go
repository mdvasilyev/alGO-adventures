package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: 	O(n)
// Space:	O(n)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = new(ListNode)
	var l1next, l2next *ListNode = l1.Next, l2.Next
	var next *ListNode = new(ListNode)
	sum := l1.Val + l2.Val
	value := sum % 10
	remainder := sum / 10
	result.Val = value
	if l1next == nil && l2next == nil {
		if remainder > 0 {
			result.Next = next
			next.Val = remainder
		}
		return result
	}
	result.Next = next
	for l1next != nil || l2next != nil || remainder != 0 {
		var l1value, l2value int
		if l1next != nil {
			l1value = l1next.Val
			l1next = l1next.Next
		} else {
			l1value = 0
		}
		if l2next != nil {
			l2value = l2next.Val
			l2next = l2next.Next
		} else {
			l2value = 0
		}
		sum = l1value + l2value + remainder
		value = sum % 10
		remainder = sum / 10
		next.Val = value
		if l1next != nil || l2next != nil || remainder > 0 {
			next.Next = new(ListNode)
			next = next.Next
		}
	}

	return result
}

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	l1.Val = 9
	l2.Val = 9
	l1.Next = new(ListNode)
	l2.Next = new(ListNode)
	l1.Next.Val = 9
	l2.Next.Val = 9
	l1.Next.Next = new(ListNode)
	l2.Next.Next = new(ListNode)
	l1.Next.Next.Val = 9
	l2.Next.Next.Val = 9
	l1.Next.Next.Next = new(ListNode)
	l2.Next.Next.Next = new(ListNode)
	l1.Next.Next.Next.Val = 9
	l2.Next.Next.Next.Val = 9
	l1.Next.Next.Next.Next = new(ListNode)
	l1.Next.Next.Next.Next.Val = 9
	l1.Next.Next.Next.Next.Next = new(ListNode)
	l1.Next.Next.Next.Next.Next.Val = 9
	l1.Next.Next.Next.Next.Next.Next = new(ListNode)
	l1.Next.Next.Next.Next.Next.Next.Val = 9

	//l1.Val = 0
	//l2.Val = 0

	//l1.Val = 2
	//l2.Val = 5
	//l1.Next = new(ListNode)
	//l2.Next = new(ListNode)
	//l1.Next.Val = 4
	//l2.Next.Val = 6
	//l1.Next.Next = new(ListNode)
	//l2.Next.Next = new(ListNode)
	//l1.Next.Next.Val = 3
	//l2.Next.Next.Val = 4

	result := addTwoNumbers(l1, l2)
	fmt.Println(result)
}
