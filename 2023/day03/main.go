package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Part struct {
	x, y, width, value int
}

func coupled(part Part, schematic []string) bool {
	for i := max(part.y-1, 0); i <= min(part.y+1, len(schematic)-1); i++ {
		for j := max(part.x-1, 0); j <= min(part.x+part.width, len(schematic[i])-1); j++ {
			if (schematic[i][j] < '0' || schematic[i][j] > '9') && schematic[i][j] != '.' {
				return true
			}
		}
	}
	return false
}

func part1(parts []Part, schematic []string) int {
	var sum int
	for _, part := range parts {
		if coupled(part, schematic) {
			sum += part.value
		}
	}
	return sum
}

func part2(parts []Part, schematic []string) int {
	var sum int
	for y := range schematic {
		for x := range schematic[y] {
			if schematic[y][x] == '*' {
				var adjacent []int
				for _, part := range parts {
					if x >= part.x-1 && x <= part.x+part.width && y >= part.y-1 && y <= part.y+1 {
						adjacent = append(adjacent, part.value)
					}
				}
				if len(adjacent) == 2 {
					sum += adjacent[0]*adjacent[1]
				}
			}
		}
	}
	return sum
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]

	var parts []Part
	for y := range lines {
		var part []byte
		var start int
		lines[y] = lines[y] + "."
		for x := range lines[y] {
			if lines[y][x] >= '0' && lines[y][x] <= '9' {
				if len(part) == 0 {
					start = x
				}
				part = append(part, lines[y][x])
			} else if len(part) > 0 {
				number, _ := strconv.Atoi(string(part))
				parts = append(parts, Part{
					x: start, y: y,
					width: x-start,
					value: number,
				})
				part = nil
			}
		}
	}

	fmt.Println(part1(parts, lines))
	fmt.Println(part2(parts, lines))
}
