package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Card struct {
	chosen, winning map[int]struct{}
}

func part1(cards []Card) int {
	var sum int
	for _, card := range cards {
		points := 0
		for number := range card.chosen {
			if _, ok := card.winning[number]; ok {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		sum += points
	}
	return sum
}

func part2(cards []Card) int {
	copies := make([]int, len(cards))

	queue := make([]int, len(cards))
	for i := range queue {
		queue[i] = i
	}

	for len(queue) > 0 {
		game := queue[0]
		copies[game]++

		var matching int
		for number := range cards[game].chosen {
			if _, ok := cards[game].winning[number]; ok {
				matching++
			}
		}

		for i := game+1; i <= game+matching && i < len(cards); i++ {
			queue = append(queue, i)
		}

		queue = queue[1:]
	}

	var sum int
	for _, count := range copies {
		sum += count
	}
	return sum
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]

	var cards []Card
	for _, line := range lines {
		id, info, _ := strings.Cut(strings.TrimSpace(strings.TrimPrefix(line, "Card")), ":")
		if id, _ := strconv.Atoi(id); id-1 != len(cards) {
			panic("Wrong game id!")
		}

		card := Card {
			chosen: make(map[int]struct{}),
			winning: make(map[int]struct{}),
		}
		set := &card.chosen
		for _, picked := range strings.Split(info, " ") {
			if picked == "|" {
				set = &card.winning
			} else if len(picked) > 0 {
				number, _ := strconv.Atoi(strings.TrimSpace(picked))
				if _, ok := (*set)[number]; ok {
					panic("Number already stored!")
				}
				(*set)[number] = struct{}{}
			}
		}

		cards = append(cards, card)
	}

	fmt.Println(part1(cards))
	fmt.Println(part2(cards))
}
