package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Edge struct {
	to     int
	weight float64
}

type PriorityQueue struct {
	vertices []VertexDistance
	indices  []int
}

type VertexDistance struct {
	vertex   int
	distance float64
}

func (pq *PriorityQueue) Len() int {
	return len(pq.vertices)
}
func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.vertices[i].distance < pq.vertices[j].distance
}
func (pq *PriorityQueue) Swap(i, j int) {
	pq.vertices[i], pq.vertices[j] = pq.vertices[j], pq.vertices[i]
	pq.indices[pq.vertices[i].vertex] = i
	pq.indices[pq.vertices[j].vertex] = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	vertex := x.(VertexDistance)
	pq.indices[vertex.vertex] = len(pq.vertices)
	pq.vertices = append(pq.vertices, vertex)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := pq.vertices
	n := len(old)
	item := old[n-1]
	pq.indices[item.vertex] = -1
	pq.vertices = old[0 : n-1]
	return item
}
func (pq *PriorityQueue) updateDistance(vertex int, distance float64) {
	index := pq.indices[vertex]
	if index != -1 {
		pq.vertices[index].distance = distance
		heap.Fix(pq, index)
	}
}

func prim(n int, adj [][]Edge) int {
	pq := &PriorityQueue{
		vertices: make([]VertexDistance, 0, n),
		indices:  make([]int, n),
	}

	for i := 0; i < n; i++ {
		pq.indices[i] = -1
		if i == 0 {
			heap.Push(pq, VertexDistance{0, 0})
		} else {
			heap.Push(pq, VertexDistance{i, 1e18})
		}
	}

	used := make([]bool, n)
	totalWeight := 0

	for pq.Len() > 0 {
		v := heap.Pop(pq).(VertexDistance)
		if used[v.vertex] {
			continue
		}

		used[v.vertex] = true
		totalWeight += int(v.distance)

		for _, e := range adj[v.vertex] {
			if !used[e.to] && e.weight < pq.vertices[pq.indices[e.to]].distance {
				pq.updateDistance(e.to, e.weight)
			}
		}
	}

	return totalWeight
}

type FastScanner struct {
	scanner *bufio.Scanner
}

func NewFastScanner() *FastScanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return &FastScanner{scanner: scanner}
}

func (fs *FastScanner) Next() int {
	fs.scanner.Scan()
	num, _ := strconv.Atoi(fs.scanner.Text())
	return num
}

func main() {
	scanner := NewFastScanner()

	n := scanner.Next()
	m := scanner.Next()

	adj := make([][]Edge, n)
	for i := range adj {
		adj[i] = make([]Edge, 0, m/n+1)
	}

	for i := 0; i < m; i++ {
		from := scanner.Next() - 1
		to := scanner.Next() - 1
		weight := scanner.Next()
		adj[from] = append(adj[from], Edge{to: to, weight: float64(weight)})
		adj[to] = append(adj[to], Edge{to: from, weight: float64(weight)})
	}

	fmt.Println(prim(n, adj))
}
