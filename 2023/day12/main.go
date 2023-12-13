package main

import (
	"os"
	"strings"
	"strconv"
	"fmt"
)

type Record struct {
	springs []byte
	groups []int
}

func (r Record) hash() string {
	return fmt.Sprint(string(r.springs), r.groups)
}

func (r Record) memoizedFill(i, group, count int, memo map[string]int) int {
	if count > 0 && (group >= len(r.groups) || count > r.groups[group]) {
		return 0
	} else if i == len(r.springs) {
		if ((group == len(r.groups)-1 && count == r.groups[group]) ||
			(group == len(r.groups) && count == 0)) {
			return 1
		}
		return 0
	}

	var arrangements int
	switch r.springs[i] {
	case '#':
		arrangements = r.memoizedFill(i+1, group, count+1, memo)
	case '.':
		if count > 0 {
			if count != r.groups[group] {
				return 0
			}
			group++
		}
		arrangements = r.memoizedFill(i+1, group, 0, memo)
	case '?':
		subproblem := Record{springs: r.springs[i:], groups: r.groups[group:]}.hash()
		if _, ok := memo[subproblem]; ok {
			return memo[subproblem]
		}

		r.springs[i] = '#'
		arrangements += r.memoizedFill(i, group, count, memo)
		r.springs[i] = '.'
		arrangements += r.memoizedFill(i, group, count, memo)
		r.springs[i] = '?'

		if count == 0 {
			memo[subproblem] = arrangements
		}
	}

	return arrangements
}

func (r Record) Fill() int {
	return r.memoizedFill(0, 0, 0, make(map[string]int))
}

func part1(records []Record) int {
	var sum int
	for _, record := range records {
		sum += record.Fill()
	}
	return sum
}

func part2(records []Record) int {
	var unfolds []Record
	for _, record := range records {
		var unfolded Record
		for i := 0; i < 5; i++ {
			unfolded.springs = append(unfolded.springs, record.springs...)
			unfolded.groups = append(unfolded.groups, record.groups...)

			if i < 4 {
				unfolded.springs = append(unfolded.springs, '?')
			}
		}
		unfolds = append(unfolds, unfolded)
	}
	return part1(unfolds)
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]

	var records []Record
	for _, line := range lines {
		springs, info, _ := strings.Cut(line, " ")

		var groups []int
		for _, group := range strings.Split(info, ",") {
			n, _ := strconv.Atoi(group)
			groups = append(groups, n)
		}

		records = append(records, Record{springs: []byte(springs), groups: groups})
	}

	fmt.Println(part1(records))
	fmt.Println(part2(records))
}
