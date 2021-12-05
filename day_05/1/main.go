package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day_05/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]line, 0)

	for scanner.Scan() {

		l := scanner.Text()
		elements := strings.Fields(l)
		start := toPoint(elements[0])
		end := toPoint(elements[2])
		lines = append(lines, line{start, end})
	}

	grid := [1000][1000]int{}
	for _, l := range lines {
		for _,p := range l.getAllPoints() {
				grid[p.x][p.y] += 1
		}
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 1 {
				result++
			}
		}
	}

	fmt.Printf("result: %d", result)


}
type point struct {
	x,y int
}

type line struct {
	start,end point
}

func (l line) minX() point {
	if l.start.x < l.end.x {
		return l.start
	}
	return l.end
}

func (l line) maxX() point {
	if l.start.x > l.end.x {
		return l.start
	}
	return l.end
}

func (l line) minY() point {
	if l.start.y < l.end.y {
		return l.start
	}
	return l.end
}

func (l line) maxY() point {
	if l.start.y > l.end.y {
		return l.start
	}
	return l.end
}

func (l line) diffY() int {
	return l.maxY().y -l.minY().y
}

func (l line) diffX() int {
	return l.maxX().x -l.minX().x
}

func (l line) getAllPoints() []point {

	points := make([]point, 0)
	if l.diffX() == 0 {
		for i := 0; i <= l.diffY(); i++ {
			points = append(points, point{l.start.x, l.minY().y + i})
		}
	} else if l.diffY() == 0{
		for i := 0; i <= l.diffX(); i++ {
			points = append(points, point{l.minX().x + i, l.end.y})
		}
	}

	return  points
}

func toPoint(s string) point{
	coordinates := strings.Split(s, ",")
	x, e1 := strconv.Atoi(coordinates[0])
	y, e2 := strconv.Atoi(coordinates[1])
	if e1 != nil || e2 != nil {
		panic("could not parse coords")
	}
	return point{x, y}
}