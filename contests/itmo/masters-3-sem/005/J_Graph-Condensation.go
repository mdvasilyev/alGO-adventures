package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph struct {
	adj       [][]int
	revAdj    [][]int
	visited   []bool
	order     []int
	comp      []int
	compNum   int
	compGraph map[int]map[int]bool
}

func newGraph(n int) *Graph {
	return &Graph{
		adj:       make([][]int, n+1),
		revAdj:    make([][]int, n+1),
		visited:   make([]bool, n+1),
		order:     make([]int, 0),
		comp:      make([]int, n+1),
		compGraph: make(map[int]map[int]bool),
	}
}

func (g *Graph) addEdge(from, to int) {
	g.adj[from] = append(g.adj[from], to)
	g.revAdj[to] = append(g.revAdj[to], from)
}

// sortingDFS выполняет первый обход в глубину для сортировки вершин
func (g *Graph) sortingDFS(v int) {
	g.visited[v] = true

	// Обходим всех соседей текущей вершины
	for _, u := range g.adj[v] {
		if !g.visited[u] {
			g.sortingDFS(u)
		}
	}

	// Добавляем вершину в конец списка
	g.order = append(g.order, v)
}

func (g *Graph) extractingDFS(v, component int) {
	g.comp[v] = component
	g.visited[v] = true

	for _, u := range g.revAdj[v] {
		if !g.visited[u] {
			g.extractingDFS(u, component)
		}
	}
}

func (g *Graph) findStrongComponents() {
	for v := 1; v < len(g.adj); v++ {
		if !g.visited[v] {
			g.sortingDFS(v)
		}
	}

	for i := range g.visited {
		g.visited[i] = false
	}

	g.compNum = 0
	for i := len(g.order) - 1; i >= 0; i-- {
		v := g.order[i]
		if !g.visited[v] {
			g.compNum++
			g.extractingDFS(v, g.compNum)
		}
	}

	for i := 1; i <= g.compNum; i++ {
		g.compGraph[i] = make(map[int]bool)
	}
}

func (g *Graph) condesationEdges() int {
	for v := 1; v < len(g.adj); v++ {
		for _, u := range g.adj[v] {
			compFrom := g.comp[v]
			compTo := g.comp[u]

			if compFrom != compTo {
				g.compGraph[compFrom][compTo] = true
			}
		}
	}

	edgeCount := 0
	for _, edges := range g.compGraph {
		edgeCount += len(edges)
	}

	return edgeCount
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	g := newGraph(n)

	for i := 0; i < m; i++ {
		scanner.Scan()
		from, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		to, _ := strconv.Atoi(scanner.Text())
		g.addEdge(from, to)
	}

	g.findStrongComponents()

	fmt.Println(g.condesationEdges())
}
