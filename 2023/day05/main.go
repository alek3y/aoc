package main

import (
	"os"
	"strings"
	"strconv"
	"fmt"
)

type Rule struct {
	to, from, size int
}

func translate(seed int, maps [][]Rule) int {
	for _, rules := range maps {
		for _, rule := range rules {
			if seed >= rule.from && seed < rule.from+rule.size {
				seed = (seed - rule.from) + rule.to
				break
			}
		}
	}
	return seed
}

func part1(seeds []int, maps [][]Rule) int {
	lowest := seeds[0]
	for _, seed := range seeds {
		lowest = min(lowest, translate(seed, maps))
	}
	return lowest
}

func invert(maps [][]Rule) [][]Rule {
	var inverted [][]Rule
	for i := range maps {
		var category []Rule
		for _, rule := range maps[len(maps)-1-i] {
			category = append(category, Rule{to: rule.from, from: rule.to, size: rule.size})
		}
		inverted = append(inverted, category)
	}
	return inverted
}

func part2(seeds []int, maps [][]Rule) int {
	inverted := invert(maps)
	for l := 0; ; l++ {
		seed := translate(l, inverted)
		for i := 0; i < len(seeds); i += 2 {
			if seed >= seeds[i] && seed <= seeds[i]+seeds[i+1] {
				return l
			}
		}
	}
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]

	var seeds []int
	for _, seed := range strings.Split(strings.TrimPrefix(lines[0], "seeds: "), " ") {
		n, _ := strconv.Atoi(seed)
		seeds = append(seeds, n)
	}

	var maps [][]Rule
	for _, line := range lines[1:] {
		if strings.Contains(line, "map:") {
			maps = append(maps, nil)
		} else if len(line) > 0 {
			var fields []int
			for _, field := range strings.Split(line, " ") {
				n, _ := strconv.Atoi(field)
				fields = append(fields, n)
			}

			category := &maps[len(maps)-1]
			*category = append(*category, Rule{to: fields[0], from: fields[1], size: fields[2]})
		}
	}

	fmt.Println(part1(seeds, maps))
	fmt.Println(part2(seeds, maps))
}
