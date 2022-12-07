package main

import (
	_ "embed"
)

//go:embed input.txt
var input string

const L = 4

func main() {
	part1()
	part2()
}

func part1() {
	runes := []rune(input)

	lastDupe := 0

	for i := 0; i < len(runes); i++ {
		for j := i - 1; j > lastDupe; j-- {
			if runes[j] == runes[i] {
				lastDupe = j
			}
		}
		if i-lastDupe > L {
			println(i)
			break
		}
	}

}

func part2() {
	// this was the same as part1 but with L = 14
}
