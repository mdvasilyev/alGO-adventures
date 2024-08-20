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
	result := new(ListNode)
	current := result
	remainder := 0
	for l1 != nil || l2 != nil || remainder != 0 {
		sum := remainder
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		remainder = sum / 10
		current.Next = new(ListNode)
		current.Next.Val = sum % 10
		current = current.Next
	}
	return result.Next
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
