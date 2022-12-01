package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	windowSize := 3
	count := 0
	depths := loadDepths()
	previous := sum(depths[0:windowSize])

	for i := 1; i <= len(depths)-windowSize; i++ {
		current := sum(depths[i : i+windowSize])

		if current > previous {
			count++
		}

		previous = current

	}

	fmt.Printf("The depth increased %d times", count)

}

func sum(window []int) (result int) {
	for _, i := range window {
		result += i
	}
	return result
}

func loadDepths() []int {

	bytes, err := ioutil.ReadFile("day_01/2/input.txt")
	if err != nil {
		panic(err)
	}

	lineStrings := strings.Split(string(bytes), "\n")
	measurements := make([]int, len(lineStrings))

	for i, v := range lineStrings {
		measurements[i], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}
	return measurements
}
