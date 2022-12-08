package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1()
	part2()
}

func part1() {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, 0)
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			up, down, left, right := true, true, true, true

			for i := col + 1; i < len(grid[row]); i++ {
				if grid[row][i] >= grid[row][col] {
					right = false
					break
				}
			}
			for i := col - 1; i >= 0; i-- {
				if grid[row][i] >= grid[row][col] {
					left = false
					break
				}
			}
			for i := row + 1; i < len(grid); i++ {
				if grid[i][col] >= grid[row][col] {
					down = false
					break
				}
			}
			for i := row - 1; i >= 0; i-- {
				if grid[i][col] >= grid[row][col] {
					up = false
					break
				}
			}
			if up || down || left || right {
				count++
			}
		}
	}
	fmt.Println(count)
}

func part2() {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, 0)
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			var up, down, left, right, score int

			for i := col + 1; i < len(grid[row]); i++ {
				if grid[row][i] >= grid[row][col] {
					right++
					break
				}
				right++
			}
			for i := col - 1; i >= 0; i-- {
				if grid[row][i] >= grid[row][col] {
					left++
					break
				}
				left++
			}
			for i := row + 1; i < len(grid); i++ {
				if grid[i][col] >= grid[row][col] {
					down++
					break
				}
				down++
			}
			for i := row - 1; i >= 0; i-- {
				if grid[i][col] >= grid[row][col] {
					up++
					break
				}
				up++
			}
			score = up * right * down * left

			if score > count {
				count = score
			}
		}
	}
	fmt.Println(count)
}
