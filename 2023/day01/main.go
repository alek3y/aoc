package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func part1(lines []string) int {
	var sum int
	for _, line := range lines {
		var first, last int
		var found bool
		for _, char := range line {
			if char >= '0' && char <= '9' {
				last = int(char - '0')
				if !found {
					first = last
					found = true
				}
			}
		}
		sum += first*10 + last
	}
	return sum
}

func part2(lines []string) int {
	digits := [...]string{
		"one", "two", "three",
		"four", "five", "six",
		"seven", "eight", "nine",
	}
	var sum int
	for _, line := range lines {
		var first, last int
		var found bool
		for i, char := range line {
			var value int
			var valid bool
			if char >= '0' && char <= '9' {
				value = int(char - '0')
				valid = true
			} else {
				for j, digit := range digits {
					if strings.HasPrefix(line[i:], digit) {
						value = j+1
						valid = true
					}
				}
			}
			if valid {
				last = value
				if !found {
					first = last
					found = true
				}
			}
		}
		sum += first*10 + last
	}
	return sum
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var lines []string
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		lines = append(lines, reader.Text())
	}

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
