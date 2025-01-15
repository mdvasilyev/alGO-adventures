package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	adj        [][]int
	visited    []bool
	inStack    []bool
	cycleStack []int
	cycle      []int
	hasCycle   bool
	cycleStart int
}

func newGraph(n int) *Graph {
	return &Graph{
		adj:        make([][]int, n),
		visited:    make([]bool, n),
		inStack:    make([]bool, n),
		cycleStack: make([]int, 0, n),
		cycle:      make([]int, 0),
		hasCycle:   false,
		cycleStart: -1,
	}
}

func (g *Graph) add(u, v int) {
	g.adj[u] = append(g.adj[u], v)
}

func (g *Graph) dfs(v int) {
	g.visited[v] = true
	g.inStack[v] = true
	g.cycleStack = append(g.cycleStack, v)

	for _, u := range g.adj[v] {
		if !g.visited[u] {
			g.dfs(u)
			if g.hasCycle {
				return
			}
		} else if g.inStack[u] {
			g.hasCycle = true
			g.cycleStart = u

			g.cycle = make([]int, 0)
			for i := len(g.cycleStack) - 1; i >= 0; i-- {
				vertex := g.cycleStack[i]
				g.cycle = append(g.cycle, vertex)
				if vertex == u {
					break
				}
			}
			for i, j := 0, len(g.cycle)-1; i < j; i, j = i+1, j-1 {
				g.cycle[i], g.cycle[j] = g.cycle[j], g.cycle[i]
			}
			return
		}
	}

	g.cycleStack = g.cycleStack[:len(g.cycleStack)-1]
	g.inStack[v] = false
}

func (g *Graph) findCycle() (bool, []int) {
	for v := 0; v < len(g.adj); v++ {
		if !g.visited[v] {
			g.dfs(v)
			if g.hasCycle {
				return true, g.cycle
			}
		}
	}
	return false, nil
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

	hasCycle, cycle := g.findCycle()

	if !hasCycle {
		fmt.Println(-1)
	} else {
		fmt.Println(len(cycle))
		for _, v := range cycle {
			fmt.Printf("%v ", v+1)
		}
	}
}
