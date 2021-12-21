package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day_14/2/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	template := ""
	steps := 10
	rules := make(map[string]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			//skip
		} else if strings.Contains(line, "->") {
			line := strings.Split(scanner.Text(), " -> ")
			rules[line[0]] = line[1]
		} else {
			template = line
		}
	}

	for i := 0; i < steps; i++ {

		fmt.Printf("Step: %d\n", i)
		pairs := make([]string, 0)
		for i := 0; i < len(template)-1; i++ {
			pairs = append(pairs, string(template[i])+string(template[i+1]))
		}

		template = ""
		for i, pair := range pairs {
			value, exists := rules[pair]
			if exists {
				if i == 0 {
					template = template + string(pair[0]) + value + string(pair[1])
				} else {
					template = template + value + string(pair[1])
				}
			}
		}
	}

	fmt.Printf("template: :%s\n", template)

	frequencies := make(map[string]int)
	for _, r := range template{
		v, exists := frequencies[string(r)]
		if exists {
			frequencies[string(r)] = v + 1
		} else {
			frequencies[string(r)] = 0
		}
	}



	fmt.Printf("frequencies: :%v\n", frequencies)
	result :=0
	fmt.Printf("result: %d ", result)

}

