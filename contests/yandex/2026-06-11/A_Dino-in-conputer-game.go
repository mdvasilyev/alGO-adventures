package main

import (
	"bufio"
	"fmt"
	"os"
)

type Obstacle struct {
	Pos int
	Type int
}

func (o Obstacle) obstacleLength() int {
	switch o.Type {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	default:
		return 0
	}
}

func (o Obstacle) obstacleScore() int {
	switch o.Type {
	case 1:
		return 1
	case 2:
		return 3
	case 3:
		return 5
	default:
		return 0
	}
}

type Jump struct {
	Pos int
	Length int
}

type Input struct {
	Obstacles []Obstacle
	Jumps []Jump
}

func getInput() *Input {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	obstacles := make([]Obstacle, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}
	for i := 0; i < n; i++ {
		obstacles[i] = Obstacle{
			Pos: a[i],
			Type:  b[i],
		}
	}

	var m int
	fmt.Fscan(in, &m)
	xs := make([]int, m)
	ys := make([]int, m)
	jumps := make([]Jump, m)

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &xs[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &ys[i])
	}
	for i := 0; i < m; i++ {
		jumps[i] = Jump{
			Pos:  xs[i],
			Length: ys[i],
		}
	}

	return &Input{
		Obstacles: obstacles,
		Jumps:     jumps,
	}
}

func checkObstacles(obstacles []Obstacle) bool {
	for i := 1; i < len(obstacles); i++ {
		prev := obstacles[i-1]
		cur := obstacles[i]

		if cur.Pos <= prev.Pos + prev.obstacleLength() {
			return false
		}
	}
	return true
}

func getScore(input *Input) int {
	res := 0
	if !checkObstacles(input.Obstacles) {
		return res
	}

	obstacles := input.Obstacles
	jumps := input.Jumps

	curJump := 0
	finalPos := 0

	for _, obstacle := range obstacles {
		for curJump < len(jumps) && jumps[curJump].Pos <= obstacle.Pos {
			pos := jumps[curJump].Pos + jumps[curJump].Length
			if pos > finalPos {
				finalPos = pos
			}
			curJump++
		}
		if finalPos >= obstacle.Pos + obstacle.obstacleLength() {
			res += obstacle.obstacleScore()
		} else {
			res--
		}
	}
	
	return max(0, res)
}

func main() {
	input := getInput()
	res := getScore(input)
	fmt.Println(res)
}
