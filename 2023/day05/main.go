package main

import (
	"os"
	"strings"
	"strconv"
	"fmt"
	"sync"
)

type Rule struct {
	to, from, size int
}

type Category struct {
	kind string
	rules []Rule
}

func translate(seed int, maps map[string]Category) int {
	kind := "seed"
	for kind != "location" {
		category := maps[kind]
		for _, rule := range category.rules {
			if seed >= rule.from && seed < rule.from+rule.size {
				seed = (seed - rule.from) + rule.to
				break
			}
		}
		kind = category.kind
	}
	return seed
}

func part1(seeds []int, maps map[string]Category) int {
	lowest := seeds[0]
	for _, seed := range seeds {
		lowest = min(lowest, translate(seed, maps))
	}
	return lowest
}

// TODO: Oh no, laziness ensued! Either do it top-down with ranges or bottom-up
func part2(seeds []int, maps map[string]Category) int {
	lowest := seeds[0]
	var lock sync.Mutex
	var wait sync.WaitGroup
	for i := 0; i < len(seeds); i += 2 {
		wait.Add(1)
		go func(i int) {
			local := lowest
			for s := seeds[i]; s < seeds[i]+seeds[i+1]; s++ {
				local = min(local, translate(s, maps))
			}

			lock.Lock()
			lowest = min(lowest, local)
			lock.Unlock()
			wait.Done()
		}(i)
	}
	wait.Wait()
	return lowest
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]

	var seeds []int
	for _, seed := range strings.Split(strings.TrimPrefix(lines[0], "seeds: "), " ") {
		n, _ := strconv.Atoi(seed)
		seeds = append(seeds, n)
	}

	var kind string
	maps := make(map[string]Category)
	for _, line := range lines[1:] {
		if strings.Contains(line, "map:") {
			header, _, _ := strings.Cut(line, " ")
			source, destination, _ := strings.Cut(header, "-to-")
			maps[source] = Category{kind: destination}
			kind = source
		} else if len(line) > 0 {
			var fields []int
			for _, field := range strings.Split(line, " ") {
				n, _ := strconv.Atoi(field)
				fields = append(fields, n)
			}

			rule := Rule{to: fields[0], from: fields[1], size: fields[2]}
			category := maps[kind]
			category.rules = append(category.rules, rule)
			maps[kind] = category
		}
	}

	fmt.Println(part1(seeds, maps))
	fmt.Println(part2(seeds, maps))
}
