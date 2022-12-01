package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	fmt.Println(calculate(newMaxComparator()))
}

func part2() {
	fmt.Println(calculate(newTop3Comparator()))
}

type comparator func(int) int

func newMaxComparator() comparator {
	max := 0
	return func(i int) int {
		if i > max {
			max = i
		}
		return max
	}
}

func newTop3Comparator() comparator {
	top3 := []int{0, 0, 0}
	return func(i int) int {
		if i > top3[0] {
			top3[0] = i
			sort.Ints(top3)
		}
		return top3[0] + top3[1] + top3[2]
	}
}

func calculate(c comparator) int {

	content, err := os.ReadFile("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	groups := strings.Split(string(content), "\n\n")
	total := 0

	for _, group := range groups {
		calories := strings.Split(group, "\n")
		sum := 0
		for _, calory := range calories {
			if calory == "" {
				continue
			}
			cal, err := strconv.Atoi(calory)
			if err != nil {
				log.Fatal(err)
			}
			sum += cal
		}
		total = c(sum)
	}
	return total
}

//
//func part1(content string) {
//
//	groups := strings.Split(content, "\n\n")
//	max := 0
//
//	for _, group := range groups {
//		calories := strings.Split(group, "\n")
//		sum := 0
//		for _, calory := range calories {
//			if calory == "" {
//				continue
//			}
//			cal, err := strconv.Atoi(calory)
//			if err != nil {
//				log.Fatal(err)
//			}
//			sum += cal
//		}
//		if sum > max {
//			max = sum
//		}
//
//	}
//	fmt.Println(max)
//}
//
//func part2(content string) {
//
//	groups := strings.Split(string(content), "\n\n")
//	top3 := []int{0, 0, 0}
//
//	for _, group := range groups {
//		calories := strings.Split(group, "\n")
//		sum := 0
//		for _, calory := range calories {
//			if calory == "" {
//				continue
//			}
//			cal, err := strconv.Atoi(calory)
//			if err != nil {
//				log.Fatal(err)
//			}
//			sum += cal
//		}
//
//		if sum > top3[0] {
//			top3[0] = sum
//			sort.Ints(top3)
//		}
//
//	}
//
//	fmt.Println(top3[0] + top3[1] + top3[2])
//}
