package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Node struct {
	val   int
	idx   int
	left  *Node
	right *Node
}

func buildBST(s []int) *Node {
	root := &Node{val: s[0], idx: 0}
	current := root
	for i := 1; i < len(s); i++ {
		current.right = &Node{val: s[i], idx: i}
		current = current.right
	}
	return root
}

func printBST(root *Node) {
	if root == nil {
		return
	}
	leftIdx, rightIdx := -1, -1
	if root.left != nil {
		leftIdx = root.left.idx + 1
	}
	if root.right != nil {
		rightIdx = root.right.idx + 1
	}
	fmt.Println(root.val, leftIdx, rightIdx)
	printBST(root.left)
	printBST(root.right)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	s := make([]int, 0, n)

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		s = append(s, value)
	}

	sort.Ints(s)

	root := buildBST(s)
	fmt.Println(n)
	printBST(root)
	fmt.Println(root.idx + 1)
}
