package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

// A* search algorithm - path finder
// https://www.geeksforgeeks.org/a-search-algorithm/
func main() {
	part1()
	part2()
}

type position struct {
	i, j int
}

func (s position) canGoTo(t position, grid [][]rune) bool {
	from := grid[s.i][s.j]
	to := grid[t.i][t.j]
	return to <= from+1
}

type node struct {
	parent *node
	position
	f    int
	g, h int
}

func part1() {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, 0)
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	end := position{}
	start := position{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'S' {
				grid[i][j] = 'a'
				start = position{i, j}
			} else if grid[i][j] == 'E' {
				grid[i][j] = 'z'
				end = position{i, j}
			}
		}
	}

	fmt.Println(findPath(grid, start, end))

}

func part2() {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, 0)

	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	end := position{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'E' {
				grid[i][j] = 'z'
				end = position{i, j}
				break
			}
		}
	}

	stepList := make([]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'a' {
				start := position{i, j}
				stepList = append(stepList, findPath(grid, start, end))
			}
		}
	}

	sort.Ints(stepList)
	fmt.Println(stepList[0])

}

func findPath(grid [][]rune, start, end position) int {

	open := make([]node, 0)
	closed := make([]node, 0)
	open = append(open, node{position: start, f: 26})

	for len(open) > 0 {
		sort.Slice(open, func(i, j int) bool {
			return open[i].f < open[j].f
		})
		q := open[0]
		open = open[1:]

		if q.position == end {
			sum := 0
			for {
				if q.parent == nil {
					fmt.Printf("got it! sum is: %d\n", sum)
					break
				}
				sum++
				q = *q.parent
			}
			return sum
		}

		neighbors := make([]node, 0)
		if q.i > 0 {
			candidate := position{q.i - 1, q.j}
			if q.position.canGoTo(candidate, grid) {
				neighbors = append(neighbors, node{parent: &q, position: candidate})
			}
		}
		if q.i < len(grid)-1 {
			candidate := position{q.i + 1, q.j}
			if q.position.canGoTo(candidate, grid) {
				neighbors = append(neighbors, node{parent: &q, position: candidate})
			}
		}
		if q.j > 0 {
			candidate := position{q.i, q.j - 1}
			if q.position.canGoTo(candidate, grid) {
				neighbors = append(neighbors, node{parent: &q, position: candidate})
			}
		}
		if q.j < len(grid[q.i])-1 {
			candidate := position{q.i, q.j + 1}
			if q.position.canGoTo(candidate, grid) {
				neighbors = append(neighbors, node{parent: &q, position: candidate})
			}
		}

		for _, n := range neighbors {
			n.h = int('z') - int(grid[n.i][n.j]) // distance between end and this node
			n.g = q.g + 1                        // distance between successor and start
			n.f = n.g + n.h

			if contains(n, open) || contains(n, closed) {
				continue
			}

			open = append(open, n)
		}

		closed = append(closed, q)

	}

	return 999_999_999
}

func contains(n node, l []node) bool {
	for _, c := range l {
		if c.position == n.position { // not sure why adding && c.f < n.f didn't work // ohhh maybe try only for open!!!!!
			return true
		}
	}
	return false
}
