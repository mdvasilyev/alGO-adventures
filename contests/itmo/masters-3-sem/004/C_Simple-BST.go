package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	val    int
	parent *Node
	left   *Node
	right  *Node
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

func insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{val: val}
	}

	current := root
	for {
		if val == current.val {
			return root
		} else if val < current.val {
			if current.left == nil {
				current.left = &Node{val: val, parent: current}
				break
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = &Node{val: val, parent: current}
				break
			}
			current = current.right
		}
	}
	return root
}

func findMin(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
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
	return root
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
