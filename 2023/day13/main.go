package main

import (
	"os"
	"strings"
	"fmt"
)

type Pattern [][]byte

func (p Pattern) Transposed() Pattern {
	var transposed Pattern
	for i := 0; i < len(p[0]); i++ {
		var column []byte
		for j := 0; j < len(p); j++ {
			column = append(column, p[j][i])
		}
		transposed = append(transposed, column)
	}
	return transposed
}

func (p Pattern) Smudges(reflection int) int {
	var smudges int
	for offset := 0; offset < min(len(p) - reflection, reflection); offset++ {
		for i := 0; i < len(p[reflection]); i++ {
			if p[reflection + offset][i] != p[reflection-1 - offset][i] {
				smudges++
			}
		}
	}
	return smudges
}

func (p Pattern) Reflection() int {
	for r := 1; r < len(p); r++ {
		if p.Smudges(r) == 0 {
			return r
		}
	}
	return 0
}

func (p Pattern) UnsmudgedReflection() int {
	for r := 1; r < len(p); r++ {
		if p.Smudges(r) == 1 {
			return r
		}
	}
	return 0
}

func part1(patterns []Pattern) int {
	var sum int
	for _, pattern := range patterns {
		sum += 100*pattern.Reflection() + pattern.Transposed().Reflection()
	}
	return sum
}

func part2(patterns []Pattern) int {
	var sum int
	for _, pattern := range patterns {
		sum += 100*pattern.UnsmudgedReflection() + pattern.Transposed().UnsmudgedReflection()
	}
	return sum
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	blocks := strings.Split(strings.TrimSuffix(string(bytes), "\n"), "\n\n")

	var patterns []Pattern
	for _, block := range blocks {
		var pattern Pattern
		for _, line := range strings.Split(block, "\n") {
			pattern = append(pattern, []byte(line))
		}
		patterns = append(patterns, pattern)
	}

	fmt.Println(part1(patterns))
	fmt.Println(part2(patterns))
}
