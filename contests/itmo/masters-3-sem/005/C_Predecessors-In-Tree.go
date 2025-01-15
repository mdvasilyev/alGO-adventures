package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	u int
	w int
}

type Graph struct {
	adj       [][]int
	visited   []bool
	entryTime []int
	exitTime  []int
	timer     int
}

func newGraph(n int) *Graph {
	return &Graph{
		adj:       make([][]int, n),
		visited:   make([]bool, n),
		entryTime: make([]int, n),
		exitTime:  make([]int, n),
	}
}

func (g *Graph) add(u, v int) {
	g.adj[u] = append(g.adj[u], v)
}

func (g *Graph) dfs(v int) {
	g.visited[v] = true
	g.timer++
	g.entryTime[v] = g.timer

	for _, u := range g.adj[v] {
		if !g.visited[u] {
			g.dfs(u)
		}
	}

	g.timer++
	g.exitTime[v] = g.timer
}

func (g *Graph) isPredecessor(u, v int) bool {
	return g.entryTime[u] <= g.entryTime[v] && g.exitTime[u] >= g.exitTime[v]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	g := newGraph(n)

	scanner.Scan()
	if n > 1 {
		parents := strings.Fields(scanner.Text())
		for i := 0; i < n-1; i++ {
			parent, _ := strconv.Atoi(parents[i])
			g.add(parent-1, i+1)
		}
	}

	pairs := make([]Pair, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		parts = strings.Fields(scanner.Text())
		u, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		pairs[i] = Pair{u: u - 1, w: w - 1}
	}

	g.dfs(0)

	for _, p := range pairs {
		if g.isPredecessor(p.u, p.w) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
