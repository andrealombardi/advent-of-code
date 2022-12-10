package main

import (
	_ "embed"
	"fmt"
	"strconv"
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

	X := 1
	C := 0
	S := 0
	cycle := func() {
		C++
		if C%40 == 20 {
			S = S + X*C
		}
	}

	for _, line := range lines {
		l := strings.Split(line, " ")

		cycle()
		if l[0] == "noop" {
			continue
		}
		cycle()
		X = X + toInt(l[1])
	}
	fmt.Println(S)

}

func part2() {
	lines := strings.Split(input, "\n")
	var X = 1
	var C = 0
	var display = ""

	draw := func() {
		if C >= X-1 && C <= X+1 {
			display += "#"
		} else {
			display += "."
		}
		C++
		if C%40 == 0 {
			fmt.Println(display)
			display = ""
			C = 0
		}
	}

	for _, line := range lines {
		l := strings.Split(line, " ")
		draw()
		if l[0] == "noop" {
			continue
		}
		draw()
		X = X + toInt(l[1])
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
