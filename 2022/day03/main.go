package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const abc = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	part1()
	part2()
}

func part1() {
	rucksacks := strings.Split(input, "\n")
	sum := 0
	for _, rucksack := range rucksacks {
		rs := []rune(rucksack)
		m1, m2 := toMap(rs[0:len(rs)/2]), toMap(rs[len(rucksack)/2:])
		d := '☕'

		for k := range m1 {
			if _, ok := m2[k]; ok {
				d = k
				sum += strings.Index(abc, string(d))
			}
		}
	}
	fmt.Println(sum)
}

func part2() {
	rucksacks := strings.Split(input, "\n")
	sum := 0
	for i := 0; i < len(rucksacks); i += 3 {

		m1, m2, m3 := toMap([]rune(rucksacks[i])), toMap([]rune(rucksacks[i+1])), toMap([]rune(rucksacks[i+2]))
		d := '☕'

		for k := range m1 {
			if _, ok := m2[k]; ok {
				if _, ok := m3[k]; ok {
					d = k
					break
				}
			}
		}
		sum += strings.Index(abc, string(d))
	}
	fmt.Println(sum)
}

func toMap(runes []rune) map[rune]struct{} {
	m := make(map[rune]struct{}, 0)
	for _, v := range runes {
		m[v] = struct{}{}
	}
	return m
}
