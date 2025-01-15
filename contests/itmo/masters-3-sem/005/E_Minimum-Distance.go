package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	adj     [][]int
	dist    []int
	visited []bool
}

func newGraph(n int) *Graph {
	g := &Graph{
		adj:     make([][]int, n),
		dist:    make([]int, n),
		visited: make([]bool, n),
	}
	for i := range g.dist {
		g.dist[i] = -1
	}
	return g
}

func (g *Graph) add(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *Graph) bfs(start int) {
	queue := make([]int, 0)
	queue = append(queue, start)
	g.visited[start] = true
	g.dist[start] = 0

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for _, u := range g.adj[v] {
			if !g.visited[u] {
				g.visited[u] = true
				g.dist[u] = g.dist[v] + 1
				queue = append(queue, u)
			}
		}
	}
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

	g.bfs(0)

	for i := 0; i < n; i++ {
		fmt.Printf("%v ", g.dist[i])
	}
}
