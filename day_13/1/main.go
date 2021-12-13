package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day_13/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]bool, 447*2+1)
	for i, _ := range grid {
		grid[i] = make([]bool, 655*2+1)
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		xy := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		grid[y][x] = true
	}

	g1 := foldX(grid, 655)
	result := 0

	for _, bools := range g1 {
		for _, b := range bools {
			if b {
				result++
			}
		}
	}

	fmt.Printf("result: %d ", result)

}

func foldX(grid [][]bool, split int) (fold [][]bool) {

	for y := 0; y < len(grid); y++ {
		line := make([]bool, 0)
		for x := 0; x < split; x++ {
			line = append(line, grid[y][x] || grid[y][len(grid[y])-(x+1)])
		}
		fold = append(fold, line)
	}

	return fold
}
