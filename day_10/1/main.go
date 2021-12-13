package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day_10/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	open := "{[(<"
	scoreMap := map[string]int {
		"}" : 0,
		"]" : 0,
		">" : 0,
		")" : 0,
	}

	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), "")
		stack := make([]string,0)
		for _, s := range elements{
			if strings.Contains(open, s) {
				stack = append(stack, s)
			} else if len(stack) == 0{
				scoreMap[s]++
				break
			} else {
				top := stack[len(stack)-1]
				if top != matching(s) {
					scoreMap[s]++
					break
				}
				stack = stack[:len(stack) -1]
			}
		}
	}


	result := scoreMap[")"] * 3 + scoreMap["]"] * 57 + scoreMap["}"] * 1197 + scoreMap[">"] * 25137


	fmt.Printf("result: %d ", result)

}

func matching(s string) string{
	switch s {
	case "}":
		return "{"
	case "]":
		return "["
	case ">":
		return "<"
	case ")":
		return "("
	default:
		panic(errors.New(s))

	}
}

