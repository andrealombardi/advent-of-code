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

var h, t point
var positions = make(map[point]int, 0)

func part2() {
	lines := strings.Split(input, "\n")
	knots := make([]point, 10)
	positions = make(map[point]int, 0)
	positions[point{0, 0}] = 0

	for _, line := range lines {
		direction := ""
		distance := 0
		fmt.Sscanf(line, "%s %d", &direction, &distance)

		for i := 0; i < distance; i++ {
			move2(direction, &knots[0])
			for j := 0; j < len(knots)-1; j++ {
				t := moveTail(knots[j], knots[j+1])
				knots[j+1] = t
			}
			positions[knots[len(knots)-1]] = 0
		}

	}
	fmt.Println(len(positions))
}

func move2(direction string, p *point) {
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

func part1() {
	lines := strings.Split(input, "\n")
	positions = make(map[point]int, 0)
	positions[t] = 0

	for _, line := range lines {
		direction := ""
		distance := 0
		fmt.Sscanf(line, "%s %d", &direction, &distance)

		for i := 0; i < distance; i++ {
			move(direction)

		}

	}
	fmt.Println(len(positions))
}

func move(direction string) {
	switch direction {
	case "R":
		if t.x < h.x {
			t = h
		}
		h = point{h.x + 1, h.y}
	case "L":
		if t.x > h.x {
			t = h
		}
		h = point{h.x - 1, h.y}
	case "U":
		if t.y < h.y {
			t = h
		}
		h = point{h.x, h.y + 1}
	case "D":
		if t.y > h.y {
			t = h
		}
		h = point{h.x, h.y - 1}
	}
	positions[t] = positions[t] + 1

}
