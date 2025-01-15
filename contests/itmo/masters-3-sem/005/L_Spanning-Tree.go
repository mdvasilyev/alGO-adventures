package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func calculateDistance(p1, p2 Point) float64 {
	dx := float64(p1.x - p2.x)
	dy := float64(p1.y - p2.y)
	return dx*dx + dy*dy
}

func prim(points []Point) float64 {
	n := len(points)
	if n <= 1 {
		return 0
	}

	minDist := make([]float64, n)
	for i := range minDist {
		minDist[i] = math.MaxFloat64
	}
	minDist[0] = 0

	used := make([]bool, n)

	totalWeight := 0.0

	for {
		v := -1
		curMin := math.MaxFloat64
		for i := 0; i < n; i++ {
			if !used[i] && minDist[i] < curMin {
				curMin = minDist[i]
				v = i
			}
		}

		if v == -1 {
			break
		}

		used[v] = true
		totalWeight += math.Sqrt(minDist[v])

		p1 := points[v]
		for to := 0; to < n; to++ {
			if !used[to] {
				dist := calculateDistance(p1, points[to])
				if dist < minDist[to] {
					minDist[to] = dist
				}
			}
		}
	}

	return totalWeight
}

func main() {
	var n int
	fmt.Scan(&n)

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&points[i].x, &points[i].y)
	}

	fmt.Printf("%.10f\n", prim(points))
}
