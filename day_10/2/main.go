package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

var scoreMultiplier = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func main() {
	f, err := os.Open("day_10/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	open := "{[(<"

	incomplete := make([][]string, 0)
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), "")
		stack := make([]string, 0)
		for i, s := range elements {
			if strings.Contains(open, s) {
				stack = append(stack, s)
			} else if len(stack) == 0 {
				break
			} else {
				top := stack[len(stack)-1]
				if top != match(s) {
					break
				}
				stack = stack[:len(stack)-1]
			}
			if i == len(elements)-1 {
				incomplete = append(incomplete, stack)
			}
		}
	}

	results := make([]int, 0)

	for i := 0; i < len(incomplete); i++ {
		sequence := make([]string, 0)
		for j := len(incomplete[i]) - 1; j >= 0; j-- {
			sequence = append(sequence, match(incomplete[i][j]))
		}
		results = append(results, scoreSequence(sequence))
	}

	sort.Ints(results)

	fmt.Printf("result: %d ", results[len(results)/2])

}

func scoreSequence(sequence []string) (score int) {
	for _, s := range sequence {
		score = score*5 + scoreMultiplier[s]
	}
	return score
}

func match(s string) string {
	switch s {
	case "}":
		return "{"
	case "]":
		return "["
	case ">":
		return "<"
	case ")":
		return "("
	case "{":
		return "}"
	case "[":
		return "]"
	case "<":
		return ">"
	case "(":
		return ")"
	default:
		panic(errors.New(s))
	}
}
