package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	vertex int
	weight int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
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

func dijkstra(graph [][]Node, source, n int) []int {
	distances := make([]int, n+1)
	for i := range distances {
		distances[i] = int(1e9)
	}
	distances[source] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Node{vertex: source, weight: 0})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Node)
		currentVertex := current.vertex

		if current.weight > distances[currentVertex] {
			continue
		}

		for _, neighbor := range graph[currentVertex] {
			newDistance := distances[currentVertex] + neighbor.weight
			if newDistance < distances[neighbor.vertex] {
				distances[neighbor.vertex] = newDistance
				heap.Push(&pq, &Node{
					vertex: neighbor.vertex,
					weight: newDistance,
				})
			}
		}
	}

	return distances
}

func main() {
	scanner := NewFastScanner()

	n := scanner.Next()
	m := scanner.Next()

	graph := make([][]Node, n+1)
	for i := range graph {
		graph[i] = make([]Node, 0)
	}

	for i := 0; i < m; i++ {
		from := scanner.Next()
		to := scanner.Next()
		weight := scanner.Next()

		graph[from] = append(graph[from], Node{vertex: to, weight: weight})
		graph[to] = append(graph[to], Node{vertex: from, weight: weight})
	}

	distances := dijkstra(graph, 1, n)

	for i := 1; i <= n; i++ {
		fmt.Printf("%v ", distances[i])
	}
}
