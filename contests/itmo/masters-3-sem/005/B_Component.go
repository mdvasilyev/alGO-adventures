package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	adj    [][]int
	colors []int
}

func (g *Graph) add(u, v int) {
	g.adj[u] = append(g.adj[u], v)
}

func (g *Graph) dfs(u, num int) {
	g.colors[u] = num
	for _, v := range g.adj[u] {
		if g.colors[v] == 0 {
			g.dfs(v, num)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	graph := &Graph{adj: make([][]int, n), colors: make([]int, n)}

	for i := 0; i < m; i++ {
		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		u, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		graph.add(u-1, w-1)
		graph.add(w-1, u-1)
	}

	number := 0

	for i := 0; i < n; i++ {
		if graph.colors[i] == 0 {
			number++
			graph.dfs(i, number)
		}
	}

	for _, col := range graph.colors {
		fmt.Printf("%v ", col)
	}
}
