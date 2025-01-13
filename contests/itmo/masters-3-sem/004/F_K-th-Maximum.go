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

type CartesianTree struct {
	root *Node
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
	} else {
		right.left = merge(left, right.left)
		updateSize(right)
		return right
	}
}

func split(root *Node, key int) (*Node, *Node) {
	if root == nil {
		return nil, nil
	}

	if root.key > key {
		left, right := split(root.left, key)
		root.left = right
		updateSize(root)
		return left, root
	} else {
		left, right := split(root.right, key)
		root.right = left
		updateSize(root)
		return root, right
	}
}

func (t *CartesianTree) Insert(key int) {
	left, right := split(t.root, key)
	t.root = merge(merge(left, newNode(key)), right)
}

func (t *CartesianTree) Delete(key int) {
	left1, right1 := split(t.root, key)
	left2, _ := split(left1, key-1)
	t.root = merge(left2, right1)
}

func (t *CartesianTree) KthMax(k int) int {
	current := t.root

	for {
		rightSize := size(current.right)

		if k == rightSize+1 {
			return current.key
		} else if k <= rightSize {
			current = current.right
		} else {
			k -= rightSize + 1
			current = current.left
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	tree := &CartesianTree{}
	var result strings.Builder

	for i := 0; i < n; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		cmd, _ := strconv.Atoi(parts[0])
		value, _ := strconv.Atoi(parts[1])

		switch cmd {
		case 1:
			tree.Insert(value)
		case 0:
			fmt.Fprintln(&result, tree.KthMax(value))
		case -1:
			tree.Delete(value)
		}
	}

	fmt.Print(result.String())
}
