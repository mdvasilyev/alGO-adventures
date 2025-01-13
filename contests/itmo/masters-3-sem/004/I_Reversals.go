package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	key      int
	priority int64
	reversed bool
	size     int
	left     *Node
	right    *Node
}

func newNode(key int, priority int64) *Node {
	return &Node{
		key:      key,
		priority: priority,
		size:     1,
	}
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.size
}

func updateSize(node *Node) {
	if node != nil {
		node.size = size(node.left) + size(node.right) + 1
	}
}

func push(root *Node) {
	if root != nil && root.reversed {
		doReversal(root.left)
		doReversal(root.right)
		root.reversed = false
	}
}

func merge(left, right *Node) *Node {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	push(left)
	push(right)

	if left.priority > right.priority {
		left.right = merge(left.right, right)
		updateSize(left)
		return left
	}
	right.left = merge(left, right.left)
	updateSize(right)
	return right
}

func split(root *Node, key int) (*Node, *Node) {
	if root == nil {
		return nil, nil
	}

	push(root)
	if size(root.left) < key {
		left, right := split(root.right, key-size(root.left)-1)
		root.right = left
		updateSize(root)
		return root, right
	}
	left, right := split(root.left, key)
	root.left = right
	updateSize(root)
	return left, root
}

func doReversal(root *Node) {
	if root != nil {
		root.reversed = !root.reversed
		root.left, root.right = root.right, root.left
	}
}

func reverse(root *Node, left, right int) *Node {
	t1, t2 := split(root, left-1)
	t3, t4 := split(t2, right-left+1)
	doReversal(t3)
	return merge(merge(t1, t3), t4)
}

func printTree(root *Node) {
	if root == nil {
		return
	}
	push(root)
	printTree(root.left)
	fmt.Printf("%d ", root.key)
	printTree(root.right)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	var root *Node

	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	for i := 1; i <= n; i++ {
		priority := rng.Int63()
		root = merge(root, newNode(i, priority))
	}

	for i := 0; i < m; i++ {
		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		root = reverse(root, left, right)
	}

	printTree(root)
}
