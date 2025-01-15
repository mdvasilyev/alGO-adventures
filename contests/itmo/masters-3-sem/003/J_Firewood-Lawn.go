package main

import (
	"fmt"
	"math"
)

var (
	Vp float64
	Vf float64
	a  float64
)

func calcDistance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

func calcTotalTime(x float64) float64 {
	fieldDist := calcDistance(0, 1, x, a)
	forestDist := calcDistance(x, a, 1, 0)

	return fieldDist/Vp + forestDist/Vf
}

func ternarySearch(left, right float64, eps float64) float64 {
	for right-left > eps {
		a := (2*left + right) / 3
		b := (left + 2*right) / 3

		if calcTotalTime(a) < calcTotalTime(b) {
			right = b
		} else {
			left = a
		}
	}
	return (left + right) / 2
}

func main() {
	var vpInt, vfInt int
	fmt.Scan(&vpInt, &vfInt)
	fmt.Scan(&a)

	Vp = float64(vpInt)
	Vf = float64(vfInt)

	result := ternarySearch(0, 1, 1e-12)

	fmt.Printf("%v\n", result)
}
