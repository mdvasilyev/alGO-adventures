package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Next *Node
	Prev *Node
	Val  int
}

type List struct {
	Head *Node
	Mid  *Node
	Len  int
}

func (l *List) PopFirst() int {
	num := -1
	if l.Len > 0 {
		num = l.Head.Next.Val
		if l.Len == 1 {
			l.Head.Next = nil
			l.Head.Prev = nil
			l.Mid = nil
		} else {
			l.Head.Next.Next.Prev = l.Head
			l.Head.Next = l.Head.Next.Next
			if l.Len%2 == 0 {
				l.Mid = l.Mid.Next
			}
		}
		l.Len--
	}
	return num
}

func (l *List) InsertBack(val int) {
	if l.Head.Next == nil && l.Head.Prev == nil {
		node := Node{Prev: l.Head, Next: l.Head, Val: val}
		l.Head.Next = &node
		l.Head.Prev = &node
		l.Mid = l.Head.Next
	} else {
		node := Node{Prev: l.Head.Prev, Next: l.Head, Val: val}
		l.Head.Prev.Next = &node
		l.Head.Prev = &node
		if l.Len%2 == 0 {
			l.Mid = l.Mid.Next
		}
	}
	l.Len++
}

func (l *List) InsertMiddle(val int) {
	if l.Head.Next == nil && l.Head.Prev == nil {
		node := Node{Prev: l.Head, Next: l.Head, Val: val}
		l.Head.Next = &node
		l.Head.Prev = &node
		l.Mid = l.Head.Next
	} else {
		node := Node{Prev: l.Mid, Next: l.Mid.Next, Val: val}
		l.Mid.Next.Prev = &node
		l.Mid.Next = &node
		if l.Len%2 == 0 {
			l.Mid = &node
		}
	}
	l.Len++
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	list := List{Head: &Node{Next: nil, Prev: nil, Val: 0}, Mid: &Node{Next: nil, Prev: nil, Val: 0}, Len: 0}

	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		cmd := strings.Split(scanner.Text(), " ")

		if cmd[0] == "+" {
			val, _ := strconv.Atoi(cmd[1])
			list.InsertBack(val)
		} else if cmd[0] == "*" {
			val, _ := strconv.Atoi(cmd[1])
			list.InsertMiddle(val)
		} else {
			fmt.Println(list.PopFirst())
		}
	}
}
