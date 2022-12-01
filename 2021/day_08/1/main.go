package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day_08/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	data := make([]io, 0)
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), "|")
		inputs := strings.Fields(elements[0])
		outputs := strings.Fields(elements[1])
		data = append(data, io{inputs, outputs})
	}

	result := 0

	for _, d := range data {
		for _, output := range d.o {
			if isOneFourSevenOrEight(output) {
				result = result + 1
			}
		}
	}

	fmt.Printf("result: %d ", result)

}

func isOneFourSevenOrEight(s string) bool{
	return len(s) == 2 || //one
	    len(s) == 4 || //four
		len(s) == 3 || //seven
		len(s) == 7 //eight
}

type io struct {
	i []string
	o []string
}
