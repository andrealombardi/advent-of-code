package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type point struct {
	x, y int
}

func main() {
	part1()
	part2()
}

func part1() {
	calculate(2)
}

func part2() {
	calculate(10)
}

func calculate(numKnots int) {
	lines := strings.Split(input, "\n")
	knots := make([]point, numKnots)
	positions := make(map[point]int, 0)
	positions[point{0, 0}] = 0

	for _, line := range lines {
		direction := ""
		distance := 0
		fmt.Sscanf(line, "%s %d", &direction, &distance)

		for i := 0; i < distance; i++ {
			moveHead(direction, &knots[0])
			for j := 0; j < len(knots)-1; j++ {
				t := moveTail(knots[j], knots[j+1])
				knots[j+1] = t
			}
			positions[knots[len(knots)-1]] = 0
		}

	}
	fmt.Println(len(positions))
}

func moveHead(direction string, p *point) {
	switch direction {
	case "R":
		p.x++
	case "L":
		p.x--
	case "U":
		p.y++
	case "D":
		p.y--
	}
}

func moveTail(head, tail point) point {
	diff := point{head.x - tail.x, head.y - tail.y}

	switch diff {
	case point{-1, 2}, point{0, 2}, point{1, 2}: //top
		return point{head.x, tail.y + 1}
	case point{-1, -2}, point{0, -2}, point{1, -2}: //bottom
		return point{head.x, tail.y - 1}
	case point{-2, 1}, point{-2, 0}, point{-2, -1}: //left
		return point{tail.x - 1, head.y}
	case point{2, 1}, point{2, 0}, point{2, -1}: //right
		return point{tail.x + 1, head.y}
	case point{2, 2}: //top-right
		return point{tail.x + 1, tail.y + 1}
	case point{-2, 2}: //top-left
		return point{tail.x - 1, tail.y + 1}
	case point{2, -2}: //bottom-right
		return point{tail.x + 1, tail.y - 1}
	case point{-2, -2}: //bottom-left
		return point{tail.x - 1, tail.y - 1}
	default:
		return tail
	}

}
