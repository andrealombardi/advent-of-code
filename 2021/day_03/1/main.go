package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day_03/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	diagnostics := make([]string, 0)

	for scanner.Scan() {
		diagnostics = append(diagnostics, scanner.Text())
	}

	gamma := ""
	numBits := len(diagnostics[0]) //take the number of bits from the first element if the list

	for i := 0 ; i < numBits; i++ {

		zeros := 0
		ones := 0

		for _, diagnostic := range diagnostics {

			digits := strings.Split(diagnostic, "")
			if digits[i] == "1" {
				ones++
			} else if digits[i] == "0" {
				zeros++
			}
		}

		if zeros > ones {
			gamma += "0"
		} else {
			gamma += "1"
		}

	}

	epsilon := not(gamma)

	gamma10, _ := strconv.ParseInt(gamma, 2, 64)
	epsilon10, _ := strconv.ParseInt(epsilon, 2, 64)


	fmt.Printf("The gamma rate is %s, in decimal is %d", gamma, gamma10)
	fmt.Printf("The epsilon rate is %s, in decimal is %d", epsilon, epsilon10)
	fmt.Printf("the product is %d", gamma10 * epsilon10)

}

//hmmm
func not(n string) (s string) {
	for _, w := range strings.Split(n, "") {
		if w == "1" {
			s += "0"
		} else {
			s += "1"
		}
	}
	fmt.Println("s is: " + s)
	return s
}