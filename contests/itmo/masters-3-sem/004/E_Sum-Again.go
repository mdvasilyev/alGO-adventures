package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type TreapNode struct {
	key      int
	priority int
	left     *TreapNode
	right    *TreapNode
	sum      int
}

type Treap struct {
	root   *TreapNode
	maxKey int
}

func newNode(key, priority int) *TreapNode {
	return &TreapNode{
		key:      key,
		priority: priority,
		sum:      key,
	}
}

func (t *Treap) getSum(node *TreapNode) int {
	if node == nil {
		return 0
	}
	return node.sum
}

func (t *Treap) updateSum(node *TreapNode) {
	if node != nil {
		node.sum = t.getSum(node.left) + t.getSum(node.right) + node.key
	}
}

func (t *Treap) split(node *TreapNode, key int) (*TreapNode, *TreapNode) {
	if node == nil {
		return nil, nil
	}

	if node.key <= key {
		right, rest := t.split(node.right, key)
		node.right = right
		t.updateSum(node)
		return node, rest
	}

	left, rest := t.split(node.left, key)
	node.left = rest
	t.updateSum(node)
	return left, node
}

func (t *Treap) merge(left, right *TreapNode) *TreapNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.priority > right.priority {
		left.right = t.merge(left.right, right)
		t.updateSum(left)
		return left
	}

	right.left = t.merge(left, right.left)
	t.updateSum(right)
	return right
}

func (t *Treap) add(key int, additionalKey *int) {
	if additionalKey != nil {
		key = (key + *additionalKey) % 1000000000
	}

	priority := rand.Intn(1000000000)
	left, right := t.split(t.root, key-1)
	middle, right := t.split(right, key)

	if middle == nil {
		middle = newNode(key, priority)
	}

	t.root = t.merge(t.merge(left, middle), right)
	if key > t.maxKey {
		t.maxKey = key
	}
}

func (t *Treap) sumRange(l, r int) int {
	if r > t.maxKey {
		r = t.maxKey
	}

	left, middle := t.split(t.root, l-1)
	middle, right := t.split(middle, r)

	sum := t.getSum(middle)
	t.root = t.merge(left, t.merge(middle, right))
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	treap := &Treap{maxKey: -1}
	var previous *int

	for i := 0; i < n; i++ {
		scanner.Scan()
		cmd := strings.Split(scanner.Text(), " ")

		if cmd[0] == "+" {
			x, _ := strconv.Atoi(cmd[1])
			treap.add(x, previous)
			previous = nil
		} else {
			l, _ := strconv.Atoi(cmd[1])
			r, _ := strconv.Atoi(cmd[2])
			sum := treap.sumRange(l, r)
			fmt.Println(sum)
			previous = &sum
		}
	}
}
