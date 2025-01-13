package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	key      int
	priority int
	size     int
	left     *Node
	right    *Node
}

func newNode(key int) *Node {
	return &Node{
		key:      key,
		priority: rand.Int(),
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

func merge(left, right *Node) *Node {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

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

func moveToFront(root *Node, left, right int) *Node {
	t1, t2 := split(root, right)
	t3, t4 := split(t1, left-1)
	return merge(merge(t4, t3), t2)
}

func printTree(root *Node) {
	if root == nil {
		return
	}
	printTree(root.left)
	fmt.Printf("%v ", root.key)
	printTree(root.right)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	var root *Node

	for i := 1; i <= n; i++ {
		root = merge(root, newNode(i))
	}

	for i := 0; i < m; i++ {
		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		root = moveToFront(root, left, right)
	}

	printTree(root)
}
