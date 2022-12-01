package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main() {
	f, err := os.Open("day_20/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()


	scanner := bufio.NewScanner(f)
	algo := ""

	image := make([]string, 0)

	for scanner.Scan() {
		line := strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), ".", "0"), "#", "1")
		if line == "" {
			//skip
		} else if len(algo) == 0{
			algo = line
		} else {
			image = append(image, line)
		}
	}

	for s := 0; s < 50; s++ {

		newImage := make([]string, 0)
		for i := -1; i < len(image) +1; i++ {
			row := ""
			for j := -1; j < len(image[0]) +1; j++ {
				index := ""
				for y := i-1; y <= i+1; y++ { //quadrant
					for x := j-1; x <= j+1; x++ {
						c := ""
						if isOut(image, x, y) { //infinite has to be reverted each iteration (had to look this one up!)
							if s%2 == 0 {
								c = string(algo[511])
							} else {
								c = string(algo[0])
							}
						} else {
							c = string(image[y][x])
						}
						index = index + c
					}
				}
				index10, err := strconv.ParseUint(index, 2, 9)
				if err != nil {
					panic(err)
				}
				row = row + string(algo[index10])
			}

			newImage = append(newImage, row)
		}

		image = newImage

	}

	result :=0
	for _, s := range image {
		for _, c := range s {
			if string(c) == "1" {
				result++
			}
		}
	}
	fmt.Printf("result: %d ", result)
}

func isOut(image []string, x,y int) bool {
	return y < 0 || y >= len(image) || x < 0 || x >= len(image[0])
}