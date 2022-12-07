package main

import (
	_ "embed"
	"fmt"
	"strings"

	. "github.com/golang-collections/collections/stack"
)

//go:embed input.txt
var input string

func main() {
	part1()
	fmt.Println()
	part2()
}

func part1() {
	inputs := strings.Split(input, "\n\n")
	stackLines, moves := strings.Split(inputs[0], "\n"), strings.Split(inputs[1], "\n")

	stacks := loadStacks(stackLines)

	for _, move := range moves {
		q, from, to := getMove(move)
		for i := 0; i < q; i++ {
			stacks[to].Push(stacks[from].Pop())
		}
	}

	for _, stack := range stacks {
		if stack.Peek() != nil {
			fmt.Print(string(stack.Peek().(rune)))
		}
	}

}

func part2() {
	inputs := strings.Split(input, "\n\n")
	stackLines, moves := strings.Split(inputs[0], "\n"), strings.Split(inputs[1], "\n")

	stacks := loadStacks(stackLines)

	for _, move := range moves {
		q, from, to := getMove(move)
		tempStack := New()
		for i := 0; i < q; i++ {
			tempStack.Push(stacks[from].Pop())
		}
		for i := 0; i < q; i++ {
			stacks[to].Push(tempStack.Pop())
		}
	}

	for _, stack := range stacks {
		if stack.Peek() != nil {
			fmt.Print(string(stack.Peek().(rune)))
		}
	}
}

func getMove(line string) (int, int, int) {
	var q, from, to int
	_, err := fmt.Sscanf(line, "move %d from %d to %d", &q, &from, &to)
	if err != nil {
		panic(err)
	}
	return q, from, to
}

// This is the trick I should have used, instead of hardcoding the columns of the stacks
//
//	for j := 0; j*4+1 < len(runes); j++ {
//		if runes[j*4+1] != ' ' {
//			stacks[j+1].Push(runes[j*4+1])
//		}
func loadStacks(stackLines []string) []Stack {
	stacks := make([]Stack, len(stackLines)+1)
	for i := len(stackLines) - 2; i >= 0; i-- {
		runes := []rune(stackLines[i])
		if runes[1] != ' ' {
			stacks[1].Push(runes[1])
		}
		if runes[5] != ' ' {
			stacks[2].Push(runes[5])
		}
		if runes[9] != ' ' {
			stacks[3].Push(runes[9])
		}
		if runes[13] != ' ' {
			stacks[4].Push(runes[13])
		}
		if runes[17] != ' ' {
			stacks[5].Push(runes[17])
		}
		if runes[21] != ' ' {
			stacks[6].Push(runes[21])
		}
		if runes[25] != ' ' {
			stacks[7].Push(runes[25])
		}
		if runes[29] != ' ' {
			stacks[8].Push(runes[29])
		}
		if runes[33] != ' ' {
			stacks[9].Push(runes[33])
		}
	}
	return stacks
}
