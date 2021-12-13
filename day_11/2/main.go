package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var grid = make([][]dumbo, 0)

func main() {
	f, err := os.Open("day_11/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), "")
		line := make([]dumbo, 0)
		for _, l := range levels {
			n, _ := strconv.Atoi(l)
			line = append(line, dumbo{n, false})
		}
		grid = append(grid, line)
	}

	result := 0
	for s := 0; true; s++ {

		if result == 100 {
			fmt.Printf("SteP: %d", s)
			break
		}

		result = 0

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				grid[i][j].increment()
			}
		}

		for {

			change := false
			for i := 0; i < len(grid); i++ {
				for j := 0; j < len(grid[i]); j++ {
					if grid[i][j].flash() {
						result++
						change = true
						neighbours(i, j)
					}
				}
			}

			if !change {
				break
			}

		}

	}
	fmt.Printf("result: %d ", result)

}

func neighbours(i, j int) {
	if i > 0 {
		grid[i-1][j].react()
		if j > 0 {
			grid[i-1][j-1].react()
		}
		if j < len(grid[i])-1 {
			grid[i-1][j+1].react()
		}
	}
	if i < len(grid)-1 {
		grid[i+1][j].react()
		if j > 0 {
			grid[i+1][j-1].react()
		}
		if j < len(grid[i])-1 {
			grid[i+1][j+1].react()
		}
	}
	if j > 0 {
		grid[i][j-1].react()
	}
	if j < len(grid[i])-1 {
		grid[i][j+1].react()
	}
}

type dumbo struct {
	energy  int
	flashed bool
}

func (d *dumbo) increment() {
	d.energy = (d.energy + 1) % 10
	d.flashed = false
}

func (d *dumbo) react() {
	if d.energy != 0 {
		d.energy = (d.energy + 1) % 10
	}
}

func (d *dumbo) flash() bool {
	if d.energy == 0 && d.flashed == false {
		d.flashed = true
		return true
	}
	return false
}
