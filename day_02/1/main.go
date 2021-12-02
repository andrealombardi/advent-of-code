package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day_02/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	coordinates := make([]string, 0)

	for scanner.Scan() {
		coordinates = append(coordinates, scanner.Text())
	}

	position := 0
	depth := 0

	for _, coordinate := range coordinates {
		pair := strings.Split(coordinate, " ")
		direction := pair[0]
		units, _ := strconv.Atoi(pair[1])

		switch direction {
		case "up":
			depth -= units
		case "down":
			depth += units
		case "forward":
			position += units
		default:
			panic(errors.New("unexpected direction"))
		}
	}


	fmt.Printf("The depth is %d, the position is %d. depth * position = %d", depth, position, depth * position)

}