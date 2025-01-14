package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Graph struct {
	numbers []int
	adj     [][]int
}

type Pair struct {
	u int
	w int
}

func adjList(n int, pairs []Pair) *Graph {
	graph := &Graph{
		numbers: make([]int, n),
		adj:     make([][]int, n),
	}

	for i := 0; i < n; i++ {
		graph.adj[i] = make([]int, 0)
	}

	for _, p := range pairs {
		graph.numbers[p.u-1]++
		graph.adj[p.u-1] = append(graph.adj[p.u-1], p.w)
	}

	for i := 0; i < n; i++ {
		if len(graph.adj[i]) > 1 {
			sort.Ints(graph.adj[i])
		}
	}

	return graph
}

func printAdjList(graph *Graph) {
	for i := 0; i < len(graph.numbers); i++ {
		fmt.Print(graph.numbers[i])
		for j := 0; j < len(graph.adj[i]); j++ {
			fmt.Printf(" %d", graph.adj[i][j])
		}
		fmt.Println()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	pairs := make([]Pair, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		u, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		pairs[i] = Pair{u: u, w: w}
	}

	fmt.Println(n)
	printAdjList(adjList(n, pairs))
}
