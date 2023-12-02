package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Color string
const (
	Red Color = "red"
	Green = "green"
	Blue = "blue"
)

func part1(games [][]map[Color]int) int {
	max := map[Color]int{Red: 12, Green: 13, Blue: 14}
	var sum int
	for i, game := range games {
		valid := true
		for _, samples := range game {
			for color, picked := range samples {
				if picked > max[color] {
					valid = false
				}
			}
		}
		if valid {
			sum += i+1
		}
	}
	return sum
}

func part2(games [][]map[Color]int) int {
	var sum int
	for _, game := range games {
		max := map[Color]int{Red: 0, Green: 0, Blue: 0}
		for _, samples := range game {
			for color, picked := range samples {
				if picked > max[color] {
					max[color] = picked
				}
			}
		}
		sum += max[Red]*max[Green]*max[Blue]
	}
	return sum
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var games [][]map[Color]int
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		id, records, _ := strings.Cut(strings.TrimPrefix(reader.Text(), "Game "), ":")
		if id, _ := strconv.Atoi(id); id-1 != len(games) {
			panic("Wrong game id!")
		}

		var game []map[Color]int
		for _, record := range strings.Split(records, ";") {
			samples := map[Color]int{Red: 0, Green: 0, Blue: 0}
			for _, picked := range strings.Split(record, ",") {
				number, color, _ := strings.Cut(strings.TrimSpace(picked), " ")
				samples[Color(color)], _ = strconv.Atoi(number)
			}
			game = append(game, samples)
		}

		games = append(games, game)
	}

	fmt.Println(part1(games))
	fmt.Println(part2(games))
}
