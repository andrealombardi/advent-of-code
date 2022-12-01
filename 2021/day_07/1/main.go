package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day_07/1/test_input.txt")
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

	sort.Ints(crabs)
	median := crabs[len(crabs)/2]
	fmt.Printf("median: %d", median)

	fuel := 0
	for i := 0; i < len(crabs); i++ {
		fuel += abs(crabs[i] - median)
	}


	fmt.Printf("total fuel needed: %d ", fuel)

}

//really
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
