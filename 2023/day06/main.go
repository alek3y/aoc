package main

import (
	"os"
	"strings"
	"strconv"
	"fmt"
)

func wins(time, distance int) int {
	var ways int
	for speed := 0; speed < time; speed++ {
		if (time - speed)*speed > distance {
			ways++
		}
	}
	return ways
}

func part1(races [2][]string) int {
	mul := 1
	for i := 0; i < len(races[0]); i++ {
		time, _ := strconv.Atoi(races[0][i])
		distance, _ := strconv.Atoi(races[1][i])
		mul *= wins(time, distance)
	}
	return mul
}

func part2(races [2][]string) int {
	time, _ := strconv.Atoi(strings.Join(races[0], ""))
	distance, _ := strconv.Atoi(strings.Join(races[1], ""))
	return wins(time, distance)
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]

	var races [2][]string
	for i, line := range lines {
		var fields []string
		_, info, _ := strings.Cut(line, ":")
		for _, field := range strings.Split(info, " ") {
			if len(field) > 0 {
				fields = append(fields, field)
			}
		}
		races[i] = fields
	}

	fmt.Println(part1(races))
	fmt.Println(part2(races))
}
