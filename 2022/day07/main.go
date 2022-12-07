package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var index int

type Node struct {
	children map[string]Node
	size     int
}

var changeDir = regexp.MustCompile("\\$ cd ([a-z]+)")
var parentDir = regexp.MustCompile("\\$ cd \\.\\.")
var file = regexp.MustCompile("(\\d+) (.+)")

func main() {
	part1()
	part2()
}

func part1() {
	index = 1
	commands := strings.Split(input, "\n")
	rootNode := parseCommands(commands)

	fmt.Println(sum1(rootNode))
}

func part2() {
	index = 1
	commands := strings.Split(input, "\n")
	rootNode := parseCommands(commands)
	candidates := sum2(rootNode, rootNode.size-(70000000-30000000))
	sort.Ints(candidates)

	fmt.Println(candidates[0])
}

func sum1(node Node) int {

	total := 0
	if node.size < 100000 && len(node.children) > 0 {
		total += node.size
	}
	for _, v := range node.children {
		total += sum1(v)
	}

	return total
}

func sum2(node Node, minSize int) []int {
	total := make([]int, 0)
	if node.size >= minSize && len(node.children) > 0 {
		total = append(total, node.size)
	}
	for _, v := range node.children {
		total = append(total, sum2(v, minSize)...)
	}

	return total
}

func parseCommands(commands []string) Node {
	children := make(map[string]Node, 0)
	for ; index < len(commands); index++ {

		line := commands[index]
		if parentDir.MatchString(line) {
			break
		} else if changeDir.MatchString(line) {
			index++
			dirName := changeDir.FindAllStringSubmatch(line, -1)[0][1]
			children[dirName] = parseCommands(commands)
		} else if file.MatchString(line) {
			fileName := file.FindAllStringSubmatch(line, -1)[0][2]
			fileSize := file.FindAllStringSubmatch(line, -1)[0][1]
			size, err := strconv.Atoi(fileSize)
			if err != nil {
				panic(err)
			}
			children[fileName] = Node{make(map[string]Node, 0), size}
		} //Don't care about dirs
	}

	size := 0
	for _, c := range children {
		size += c.size
	}

	return Node{children: children, size: size}
}
