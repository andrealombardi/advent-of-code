package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day_03/2/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	diagnostics := make([]string, 0)

	for scanner.Scan() {
		diagnostics = append(diagnostics, scanner.Text())
	}

	o2 := calculateRating(diagnostics, "o2")
	co2 := calculateRating(diagnostics, "co2")

	fmt.Printf("the product is %d", o2*co2)

}

func calculateRating(candidates []string, gas string) int64 {
	numBits := len(candidates[0])
	for i := 0; i < numBits; i++ {
		candidates = reduce(candidates, i, gas)
		if len(candidates) == 1 {
			gasBase2 := candidates[0]
			gasBase10, _ := strconv.ParseInt(gasBase2, 2, 64)
			fmt.Printf("The %s rating is %s, in decimal is %d", gas, gasBase2, gasBase10)
			return gasBase10
		}
	}
	panic("could not find candidate")
}

func reduce(candidates []string, i int, gas string) []string {
	zeros := make([]string, 0)
	ones := make([]string, 0)

	for _, candidate := range candidates {
		digits := strings.Split(candidate, "")
		if digits[i] == "1" {
			ones = append(ones, candidate)
		} else if digits[i] == "0" {
			zeros = append(zeros, candidate)
		}
	}

	switch gas {
	case "o2":
		if len(zeros) > len(ones) {
			return zeros
		} else {
			return ones
		}
	case "co2":
		if len(zeros) > len(ones) {
			return ones
		} else {
			return zeros
		}
	default:
		panic("what gas is this?")
	}

}
