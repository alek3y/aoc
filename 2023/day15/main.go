package main

import (
	"os"
	"strings"
	"strconv"
	"fmt"
)

func hash(step string) int {
	value := 0
	for _, char := range step {
		value += int(char)
		value *= 17
		value %= 256
	}
	return value
}

func part1(sequence []string) int {
	var sum int
	for _, step := range sequence {
		sum += hash(step)
	}
	return sum
}

type Lens struct {
	label string
	length int
}

func part2(sequence []string) int {
	var boxes [256][]Lens
	for _, step := range sequence {
		label, length, assign := strings.Cut(strings.ReplaceAll(step, "-", ""), "=")
		box := hash(label)

		lens := -1
		for i := 0; i < len(boxes[box]); i++ {
			if label == boxes[box][i].label {
				lens = i
				break
			}
		}

		if assign {
			n, _ := strconv.Atoi(length)
			if lens >= 0 {
				boxes[box][lens].length = n
			} else {
				boxes[box] = append(boxes[box], Lens{label: label, length: n})
			}
		} else if lens >= 0 {
			boxes[box] = append(boxes[box][:lens], boxes[box][lens+1:]...)
		}
	}

	var sum int
	for i := range boxes {
		for j := range boxes[i] {
			sum += (i+1) * (j+1) * boxes[i][j].length
		}
	}
	return sum
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	sequence := strings.Split(strings.ReplaceAll(string(bytes), "\n", ""), ",")

	fmt.Println(part1(sequence))
	fmt.Println(part2(sequence))
}
