package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("day_01/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	count := -1 //ignore the first input
	previous := int64(0)
	for scanner.Scan() {
		current, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		if current > previous {
			count++
		}
		previous = current
	}

	fmt.Printf("The depth increased %d times", count)

}
