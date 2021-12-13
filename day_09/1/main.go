package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day_09/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := make([][]int, 0)
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), "")
		row := make([]int,0)
		for _, s := range elements{
			point,_ := strconv.Atoi(s)
			row = append(row, point)
		}
		grid = append(grid, row)
	}

	result := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isHorizontalMin(j, grid[i]) && isVerticalMin(i,j, grid){
				fmt.Printf("[%d,%d]", i,j)
				result += grid[i][j] + 1
			}
		}
	}

	fmt.Printf("result: %d ", result)

}

func isHorizontalMin(j int, row []int)  bool{
	min := true
	if j > 0 {
		min = min && row[j] < row[j-1]
	}
	if j < len(row) -1 {
		min = min && row[j] < row[j+1]
	}
	return min
}

func isVerticalMin(i, j int, grid [][]int) bool {
	min := true
	if i > 0 {
		min = min && grid[i][j] < grid[i-1][j]
	}
	if i < len(grid) -1 {
		min = min && grid[i][j] < grid[i+1][j]
	}
	return min
}

