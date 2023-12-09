package main

import (
	"os"
	"strings"
	"strconv"
	"fmt"
)

func predict(history []int) int {
	zeroes := true
	for _, value := range history {
		if value != 0 {
			zeroes = false
		}
	}
	if zeroes {
		return 0
	}

	var differences []int
	for i := 1; i < len(history); i++ {
		difference := history[i] - history[i-1]
		differences = append(differences, difference)
	}

	return history[len(history)-1] + predict(differences)
}

func part1(histories [][]int) int {
	var sum int
	for _, history := range histories {
		sum += predict(history)
	}
	return sum
}

func part2(histories [][]int) int {
	var sum int
	for _, history := range histories {
		var reversed []int
		for i := range history {
			reversed = append(reversed, history[len(history)-1-i])
		}
		sum += predict(reversed)
	}
	return sum
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]

	var histories [][]int
	for _, line := range lines {
		var history []int
		for _, value := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(value)
			history = append(history, n)
		}
		histories = append(histories, history)
	}

	fmt.Println(part1(histories))
	fmt.Println(part2(histories))
}
