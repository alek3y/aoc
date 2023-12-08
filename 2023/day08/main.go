package main

import (
	"os"
	"strings"
	"fmt"
)

type Pair[T, S any] struct {
	first T
	second S
}

func part1(directions string, network map[string]Pair[string, string]) int {
	var steps int
	node := "AAA"
	for ; node != "ZZZ"; steps++ {
		direction := directions[steps % len(directions)]
		switch direction {
		case 'L':
			node = network[node].first
		case 'R':
			node = network[node].second
		}
	}
	return steps
}

func part2(directions string, network map[string]Pair[string, string]) int {
	var paths []Pair[int, int]
	for node := range network {
		if node[len(node)-1] != 'A' {
			continue
		}

		var steps int
		var ancestors []string
		cycle := -1
		it := node
		for cycle < 0 {	// Assumes every node runs into cycles
			direction := steps % len(directions)

			for i := range ancestors {
				if ancestors[i] == it && i % len(directions) == direction {
					cycle = i
					break
				}
			}

			if cycle < 0 {
				ancestors = append(ancestors, it)

				switch directions[direction] {
				case 'L':
					it = network[it].first
				case 'R':
					it = network[it].second
				}

				steps++
			}
		}

		destination := -1
		for i := range ancestors {
			if ancestors[i][len(ancestors[i])-1] == 'Z' {
				destination = i	// Assumes the last 'Z' node is the destination
			}
		}

		paths = append(paths, Pair[int, int]{
			first: destination,
			second: len(ancestors)-cycle,
		})
	}

	offsets := make([]int, len(paths))
	for i := range offsets {
		offsets[i] = paths[i].first
	}

	var aligned bool
	for !aligned {
		var biggest int
		for i := range offsets {
			if offsets[i] > offsets[biggest] {
				biggest = i
			}
		}

		for i := range offsets {
			if offsets[i] < offsets[biggest] {
				offsets[i] += paths[i].second
			}
		}

		aligned = true
		for i := 1; i < len(offsets); i++ {
			if offsets[i] != offsets[i-1] {
				aligned = false
				break
			}
		}
	}

	return offsets[0]
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]

	directions := lines[0]
	network := make(map[string]Pair[string, string])
	for _, line := range lines[1:] {
		if len(line) > 0 {
			from, to, _ := strings.Cut(line, " = ")
			left, right, _ := strings.Cut(strings.Trim(to, "()"), ", ")
			network[from] = Pair[string, string]{first: left, second: right}
		}
	}

	fmt.Println(part1(directions, network))
	fmt.Println(part2(directions, network))
}
