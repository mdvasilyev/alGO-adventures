package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	adj      [][]int
	visited  []bool
	inStack  []bool
	sorted   []int
	hasCycle bool
}

func newGraph(n int) *Graph {
	return &Graph{
		adj:      make([][]int, n),
		visited:  make([]bool, n),
		inStack:  make([]bool, n),
		sorted:   make([]int, 0, n),
		hasCycle: false,
	}
}

func (g *Graph) add(u, v int) {
	g.adj[u] = append(g.adj[u], v)
}

func (g *Graph) dfs(v int) {
	g.visited[v] = true
	g.inStack[v] = true

	for _, u := range g.adj[v] {
		if !g.visited[u] {
			g.dfs(u)
		} else if g.inStack[u] {
			g.hasCycle = true
		}
		if g.hasCycle {
			return
		}
	}

	g.inStack[v] = false
	g.sorted = append(g.sorted, v)
	copy(g.sorted[1:], g.sorted)
	g.sorted[0] = v
}

func (g *Graph) topSort() ([]int, bool) {
	for v := 0; v < len(g.adj); v++ {
		if !g.visited[v] {
			g.dfs(v)
			if g.hasCycle {
				return nil, false
			}
		}
	}

	return g.sorted, true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	g := newGraph(n)

	for i := 0; i < m; i++ {
		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		g.add(u-1, v-1)
	}

	sorted, possible := g.topSort()

	if !possible {
		fmt.Println(-1)
	} else {
		for _, v := range sorted {
			fmt.Printf("%v ", v+1)
		}
	}
}
