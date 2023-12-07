package main

import (
	"os"
	"strings"
	"fmt"
	"strconv"
	"sort"
)

type Hand struct {
	cards string
	bid, kind int
}

type Game struct {
	hands []Hand
	strengths string
}

func kind(cards string) int {
	counts := make(map[rune]int)
	for _, card := range cards {
		count := strings.Count(cards, string(card))
		counts[card] = count
	}

	var pairs [2]int
	i := 0
	for _, count := range counts {
		if count > 1 {
			pairs[i] = count
			i++
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(pairs[:])))

	return pairs[0]*10 + pairs[1]
}

func (g Game) Strength(card byte) int {
	return strings.Index(g.strengths, string(card))
}

func (g Game) Len() int {
	return len(g.hands)
}

func (g Game) Swap(i, j int) {
	g.hands[i], g.hands[j] = g.hands[j], g.hands[i]
}

func (g Game) Less(i, j int) bool {
	var idxLeft, idxRight int
	for k := 0; k < len(g.hands[i].cards) && idxLeft == idxRight; k++ {
		idxLeft = g.Strength(g.hands[i].cards[k])
		idxRight = g.Strength(g.hands[j].cards[k])
	}

	a, b := g.hands[i].kind, g.hands[j].kind
	return a < b || (a == b && idxLeft < idxRight)
}

func part1(game Game) int {
	var sum int
	sort.Sort(game)
	for i, hand := range game.hands {
		sum += (i+1) * hand.bid
	}
	return sum
}

func part2(game Game) int {
	for i, hand := range game.hands {
		biggest := 'J'
		var countBiggest int
		for _, card := range hand.cards {
			countCard := strings.Count(hand.cards, string(card))
			if biggest == 'J' || (card != 'J' && countCard > countBiggest) {
				biggest = card
				countBiggest = countCard
			}
		}
		game.hands[i].kind = kind(strings.ReplaceAll(hand.cards, "J", string(biggest)))
	}
	game.strengths = "J" + strings.ReplaceAll(game.strengths, "J", "")
	return part1(game)
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]

	game := Game{strengths: "23456789TJQKA"}
	for _, line := range lines {
		cards, bid, _ := strings.Cut(line, " ")
		n, _ := strconv.Atoi(bid)
		game.hands = append(game.hands, Hand{
			cards: cards, bid: n,
			kind: kind(cards),
		})
	}

	fmt.Println(part1(game))
	fmt.Println(part2(game))
}
