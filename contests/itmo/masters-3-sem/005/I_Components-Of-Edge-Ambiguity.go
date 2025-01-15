package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Edge struct {
	to      int
	edgeNum int
	isMult  bool
}

type Graph struct {
	adj      [][]Edge
	used     []int
	enter    []int
	ret      []int
	colors   []int
	timer    int
	maxColor int
	bridges  map[int]bool
	edgeNums map[int]map[int]int
}

func newGraph(n int) *Graph {
	g := &Graph{
		adj:      make([][]Edge, n+1),
		used:     make([]int, n+1),
		enter:    make([]int, n+1),
		ret:      make([]int, n+1),
		colors:   make([]int, n+1),
		timer:    1,
		bridges:  make(map[int]bool),
		edgeNums: make(map[int]map[int]int),
	}

	for i := 1; i <= n; i++ {
		g.edgeNums[i] = make(map[int]int)
	}
	return g
}

func (g *Graph) getEdgeNumber(v, u int) int {
	if v > u {
		v, u = u, v
	}
	return g.edgeNums[v][u]
}

func (g *Graph) addEdge(v, u, edgeNum int) {
	if v > u {
		v, u = u, v
	}

	isMult := false
	if prevNum := g.edgeNums[v][u]; prevNum != 0 {
		isMult = true
		g.updateEdgeMult(v, u, prevNum)
	}

	g.edgeNums[v][u] = edgeNum
	g.adj[v] = append(g.adj[v], Edge{to: u, edgeNum: edgeNum, isMult: isMult})
	g.adj[u] = append(g.adj[u], Edge{to: v, edgeNum: edgeNum, isMult: isMult})
}

func (g *Graph) updateEdgeMult(v, u, edgeNum int) {
	for i := range g.adj[v] {
		if g.adj[v][i].to == u && g.adj[v][i].edgeNum == edgeNum {
			g.adj[v][i].isMult = true
		}
	}
	for i := range g.adj[u] {
		if g.adj[u][i].to == v && g.adj[u][i].edgeNum == edgeNum {
			g.adj[u][i].isMult = true
		}
	}
}

func (g *Graph) dfs(v, parent int) {
	g.used[v] = 1
	g.enter[v] = g.timer
	g.ret[v] = g.timer
	g.timer++

	for _, e := range g.adj[v] {
		if e.to == parent {
			continue
		}

		if g.used[e.to] == 1 {
			g.ret[v] = min(g.ret[v], g.enter[e.to])
		} else {
			g.dfs(e.to, v)
			g.ret[v] = min(g.ret[v], g.ret[e.to])
			if g.ret[e.to] > g.enter[v] && !e.isMult {
				g.bridges[e.edgeNum] = true
			}
		}
	}
}

func (g *Graph) setColors(v, color int) {
	g.colors[v] = color

	for _, e := range g.adj[v] {
		if g.colors[e.to] == 0 {
			if g.bridges[e.edgeNum] {
				g.maxColor++
				g.setColors(e.to, g.maxColor)
			} else {
				g.setColors(e.to, color)
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	g := newGraph(n)

	for i := 1; i <= m; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		u, _ := strconv.Atoi(scanner.Text())
		g.addEdge(v, u, i)
	}

	for v := 1; v < len(g.adj); v++ {
		if g.used[v] == 0 {
			g.dfs(v, -1)
		}
	}

	g.maxColor = 0
	for v := 1; v < len(g.adj); v++ {
		if g.colors[v] == 0 {
			g.maxColor++
			g.setColors(v, g.maxColor)
		}
	}

	fmt.Println(g.maxColor)
	for i := 1; i <= n; i++ {
		fmt.Printf("%v ", g.colors[i])
	}
}
