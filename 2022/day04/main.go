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
	count := 0
	for _, line := range lines {
		startX, endX, startY, endY := extractRanges(line)
		if (startX >= startY && endX <= endY) || (startY >= startX && endY <= endX) {
			count++
		}
	}
	fmt.Println(count)
}

func part2() {
	lines := strings.Split(input, "\n")
	count := 0
	for _, line := range lines {
		startX, endX, startY, endY := extractRanges(line)
		if startX <= endY && endX >= startY {
			count++
		}
	}
	fmt.Println(count)
}

func extractRanges(line string) (int, int, int, int) {
	var startX, endX, startY, endY int
	_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &startX, &endX, &startY, &endY)
	if err != nil {
		panic(err)
	}
	return startX, endX, startY, endY
}
