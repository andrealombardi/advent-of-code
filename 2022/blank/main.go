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
	fmt.Println(lines)

}

func part2() {
}

