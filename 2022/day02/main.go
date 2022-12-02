package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	strategy := make([][]string, 0)
	for _, line := range strings.Split(string(content), "\n") {
		strategy = append(strategy, strings.Split(line, " "))
	}

	part1(strategy)
	part2(strategy)
}

func part1(strategy [][]string) {
	score := 0
	for _, game := range strategy {
		score += calculateScore(game[0], game[1])
	}
	fmt.Println(score)
}

func part2(strategy [][]string) {
	score := 0
	for _, game := range strategy {
		score += calculateOutcome(game[0], game[1])
	}
	fmt.Println(score)
}

func calculateScore(s string, t string) int {

	switch true {
	// rock
	case s == "A" && t == "X":
		return 1 + 3
	case s == "A" && t == "Y":
		return 2 + 6
	case s == "A" && t == "Z":
		return 3 + 0
	// paper
	case s == "B" && t == "X":
		return 1 + 0
	case s == "B" && t == "Y":
		return 2 + 3
	case s == "B" && t == "Z":
		return 3 + 6
	// scissor
	case s == "C" && t == "X":
		return 1 + 6
	case s == "C" && t == "Y":
		return 2 + 0
	case s == "C" && t == "Z":
		return 3 + 3
	// ?
	default:
		log.Fatal(s, t)
	}
	return 0
}

func calculateOutcome(s string, t string) int {

	switch true {
	case s == "A" && t == "X":
		return calculateScore(s, "Z")
	case s == "A" && t == "Y":
		return calculateScore(s, "X")
	case s == "A" && t == "Z":
		return calculateScore(s, "Y")
	case s == "B" && t == "X":
		return calculateScore(s, "X")
	case s == "B" && t == "Y":
		return calculateScore(s, "Y")
	case s == "B" && t == "Z":
		return calculateScore(s, "Z")
	case s == "C" && t == "X":
		return calculateScore(s, "Y")
	case s == "C" && t == "Y":
		return calculateScore(s, "Z")
	case s == "C" && t == "Z":
		return calculateScore(s, "X")
	default:
		log.Fatal(s, t)
	}
	return 0
}
