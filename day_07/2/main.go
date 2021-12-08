package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day_07/2/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	crabs := make([]int, 0)
	for scanner.Scan() {

		elements := strings.Split(scanner.Text(), ",")
		for _, element := range elements {
			position, err := strconv.Atoi(element)
			if err != nil {
				panic(err)
			}
			crabs = append(crabs, position)
		}

	}

	avg := 0
	for _, crab := range crabs {
		avg += crab
	}


	average := int(float64(avg)/float64(len(crabs)))
	fmt.Printf("Average: %d", average)


	fuel := 0
	for i := 0; i < len(crabs); i++ {
		steps := abs(crabs[i] - average)
		for i := 1; i <= steps; i++ {
			fuel += i
		}
	}

	fmt.Printf("total fuel needed: %d ", fuel)

}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
