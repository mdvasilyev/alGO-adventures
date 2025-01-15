package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Edge struct {
	from int
	to   int
	id   int
}

type Graph struct {
	adj     [][]Edge
	visited []bool
	inTime  []int
	minTime []int
	timer   int
	bridges []int
}

func newGraph(n int) *Graph {
	return &Graph{
		adj:     make([][]Edge, n),
		visited: make([]bool, n),
		inTime:  make([]int, n),
		minTime: make([]int, n),
		timer:   0,
		bridges: make([]int, 0),
	}
}

func (g *Graph) add(u, v, id int) {
	g.adj[u] = append(g.adj[u], Edge{from: u, to: v, id: id})
	g.adj[v] = append(g.adj[v], Edge{from: v, to: u, id: id})
}

func (g *Graph) dfs(v, parent int) {
	g.visited[v] = true
	g.inTime[v] = g.timer
	g.minTime[v] = g.timer
	g.timer++

	for _, e := range g.adj[v] {
		to := e.to
		if to == parent {
			continue
		}
		if g.visited[to] {
			if g.inTime[to] < g.minTime[v] {
				g.minTime[v] = g.inTime[to]
			}
		} else {
			g.dfs(to, v)
			if g.minTime[to] < g.minTime[v] {
				g.minTime[v] = g.minTime[to]
			}
			if g.minTime[to] > g.inTime[v] {
				g.bridges = append(g.bridges, e.id)
			}
		}
	}
}

func (g *Graph) findBridges() []int {
	for v := 0; v < len(g.adj); v++ {
		if !g.visited[v] {
			g.dfs(v, -1)
		}
	}

	sort.Ints(g.bridges)
	return g.bridges
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
		g.add(u-1, v-1, i+1)
	}

	bridges := g.findBridges()

	fmt.Println(len(bridges))
	if len(bridges) > 0 {
		for _, id := range bridges {
			fmt.Printf("%v ", id)
		}
	}
}
