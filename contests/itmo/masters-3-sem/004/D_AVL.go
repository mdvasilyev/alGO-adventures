package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	val    int
	h      int
	parent *Node
	left   *Node
	right  *Node
}

func (root *Node) height() int {
	if root == nil {
		return 0
	}
	return root.h
}

func (root *Node) updateHeight() {
	root.h = 1 + max(root.left.height(), root.right.height())
}

func (root *Node) heightDiff() int {
	return root.left.height() - root.right.height()
}

func rotateRight(y *Node) *Node {
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	x.parent = y.parent
	y.parent = x
	if T2 != nil {
		T2.parent = y
	}

	y.updateHeight()
	x.updateHeight()

	return x
}

func rotateLeft(x *Node) *Node {
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	y.parent = x.parent
	x.parent = y
	if T2 != nil {
		T2.parent = x
	}

	x.updateHeight()
	y.updateHeight()

	return y
}

func balance(node *Node) *Node {
	if node == nil {
		return nil
	}

	node.updateHeight()
	balanceFactor := node.heightDiff()

	if balanceFactor > 1 && node.left.heightDiff() < 0 {
		node.left = rotateLeft(node.left)
		return rotateRight(node)
	}

	if balanceFactor < -1 && node.right.heightDiff() > 0 {
		node.right = rotateRight(node.right)
		return rotateLeft(node)
	}

	if balanceFactor > 1 {
		return rotateRight(node)
	}

	if balanceFactor < -1 {
		return rotateLeft(node)
	}

	return node
}

func insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{val: val, h: 1}
	}

	if val < root.val {
		root.left = insert(root.left, val)
		root.left.parent = root
	} else if val > root.val {
		root.right = insert(root.right, val)
		root.right.parent = root
	} else {
		return root
	}

	return balance(root)
}

func findMin(node *Node) *Node {
	current := node
	for current != nil && current.left != nil {
		current = current.left
	}
	return current
}

func delete(root *Node, val int) *Node {
	if root == nil {
		return nil
	}

	if val < root.val {
		root.left = delete(root.left, val)
		if root.left != nil {
			root.left.parent = root
		}
	} else if val > root.val {
		root.right = delete(root.right, val)
		if root.right != nil {
			root.right.parent = root
		}
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		successor := findMin(root.right)
		root.val = successor.val
		root.right = delete(root.right, successor.val)
		if root.right != nil {
			root.right.parent = root
		}
	}

	return balance(root)
}

func (root *Node) exists(val int) {
	if root == nil {
		fmt.Println("false")
		return
	}
	if root.val == val {
		fmt.Println("true")
	} else if root.val > val {
		root.left.exists(val)
	} else {
		root.right.exists(val)
	}
}

func (root *Node) next(val int) {
	var successor *Node
	current := root

	for current != nil {
		if current.val > val {
			successor = current
			current = current.left
		} else {
			current = current.right
		}
	}

	if successor == nil {
		fmt.Println("none")
	} else {
		fmt.Println(successor.val)
	}
}

func (root *Node) prev(val int) {
	var predecessor *Node
	current := root

	for current != nil {
		if current.val < val {
			predecessor = current
			current = current.right
		} else {
			current = current.left
		}
	}

	if predecessor == nil {
		fmt.Println("none")
	} else {
		fmt.Println(predecessor.val)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var root *Node

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		command := parts[0]
		value := 0
		if len(parts) > 1 {
			fmt.Sscanf(parts[1], "%d", &value)
		}

		switch command {
		case "insert":
			root = insert(root, value)
		case "delete":
			root = delete(root, value)
		case "exists":
			if root == nil {
				fmt.Println("false")
			} else {
				root.exists(value)
			}
		case "next":
			if root == nil {
				fmt.Println("none")
			} else {
				root.next(value)
			}
		case "prev":
			if root == nil {
				fmt.Println("none")
			} else {
				root.prev(value)
			}
		}
	}
}
