package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Heap struct {
	s []int
}

func (h *Heap) Insert(val int) {
	h.s = append(h.s, val)
	for idx := len(h.s) - 1; h.s[idx] > h.s[(idx-1)/2]; idx = (idx - 1) / 2 {
		h.s[idx], h.s[(idx-1)/2] = h.s[(idx-1)/2], h.s[idx]
	}
}

func (h *Heap) Extract() {
	heapLen := len(h.s)
	fmt.Println(h.s[0])
	h.s[0], h.s[heapLen-1] = h.s[heapLen-1], h.s[0]
	h.s = h.s[:heapLen-1]
	heapLen--
	for idx := 0; 2*idx+1 < heapLen; {
		leftChildIdx, rightChildIdx := 2*idx+1, 2*idx+2
		newIdx := leftChildIdx
		if rightChildIdx < heapLen && h.s[rightChildIdx] > h.s[leftChildIdx] {
			newIdx = rightChildIdx
		}
		if h.s[idx] >= h.s[newIdx] {
			break
		}
		h.s[idx], h.s[newIdx] = h.s[newIdx], h.s[idx]
		idx = newIdx
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	h := &Heap{}

	for i := 0; i < n; i++ {
		scanner.Scan()
		cmd := scanner.Text()
		if cmd == "0" {
			scanner.Scan()
			val, _ := strconv.Atoi(scanner.Text())
			h.Insert(val)
		} else {
			h.Extract()
		}
	}
}
