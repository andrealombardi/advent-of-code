package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var state = map[string]int{
	"w": 0,
	"x": 0,
	"y": 0,
	"z": 0,
}

func main() {
	f, err := os.Open("day_24/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	instructions := make([][]string, 0)

	for scanner.Scan() {
		instructions = append(instructions, strings.Split(scanner.Text(), " "))
	}

	candidate := ""
	for i := 99999999999999; i > 0; i-- {

		candidate = fmt.Sprint(i)
		if strings.Contains(candidate, "0") {
			continue
		}
		index := 0
		for _, instruction := range instructions {
			switch instruction[0] {
			case "inp":
				v, _ := strconv.Atoi(string(candidate[index]))
				index++
				inp(instruction[1], v)
				break
			case "add":
				add(instruction[1], toInt(instruction[2]))
				break
			case "mul":
				mul(instruction[1], toInt(instruction[2]))
				break
			case "div":
				div(instruction[1], toInt(instruction[2]))
				break
			case "mod":
				mod(instruction[1], toInt(instruction[2]))
				break
			case "eql":
				eql(instruction[1], toInt(instruction[2]))
				break
			default:
				panic(errors.New("Unknow instruction: " + instruction[0]))
			}
		}

		if state["z"] == 0 {
			break
		}
			
	}



	fmt.Printf("result: %s\n", candidate)

}

func inp(destination string, value int) {
	state[destination] = value
}

func add(a string, b int) {
	inp(a, state[a]+b)
}

func div(a string, b int) {
	inp(a, state[a]/b)
}

func mul(a string, b int) {
	inp(a, state[a]*b)
}

func mod(a string, b int) {
	inp(a, state[a]%b)
}

func eql(a string, b int) {
	if state[a] == b {
		inp(a, 1)
	} else {
		inp(a, 0)
	}
}

func dump(desc string, index int) {
	fmt.Printf("%s, %d --> ", desc, index)
	for k, v := range state {
		fmt.Printf("%s: %d, ", k, v)
	}
	fmt.Println()
}

func toInt(s string) int {
	v, exists := state[s]
	if exists {
		return v
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
