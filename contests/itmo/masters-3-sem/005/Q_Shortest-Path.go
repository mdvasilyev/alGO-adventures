package main

import (
	"fmt"
	"math"
)

func findShortestPath(n, source, target int, graph [][]int64) int64 {
	dist := make([]int64, n)
	visited := make([]bool, n)

	for i := range dist {
		dist[i] = math.MaxInt64
	}

	dist[source] = 0

	for i := 0; i < n; i++ {
		minDist := int64(math.MaxInt64)
		current := -1
		for v := 0; v < n; v++ {
			if !visited[v] && dist[v] < minDist {
				minDist = dist[v]
				current = v
			}
		}

		if current == -1 {
			break
		}

		visited[current] = true

		for next := 0; next < n; next++ {
			if graph[current][next] == -1 || visited[next] {
				continue
			}

			newDist := dist[current]
			if newDist < math.MaxInt64-graph[current][next] {
				newDist += graph[current][next]
				if newDist < dist[next] {
					dist[next] = newDist
				}
			}
		}
		if current == target {
			break
		}
	}

	if dist[target] == math.MaxInt64 {
		return -1
	}
	return dist[target]
}

func main() {
	var n, s, f int
	fmt.Scan(&n, &s, &f)

	graph := make([][]int64, n)
	for i := range graph {
		graph[i] = make([]int64, n)
		for j := range graph[i] {
			fmt.Scan(&graph[i][j])
		}
	}

	fmt.Println(findShortestPath(n, s-1, f-1, graph))
}
