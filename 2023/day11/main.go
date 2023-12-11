package main

import (
	"os"
	"strings"
	"fmt"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Coord struct {
	x, y int
}

func (lhs Coord) Distance(rhs Coord) int {
	return Abs(rhs.x - lhs.x) + Abs(rhs.y - lhs.y)
}

func distances(galaxies []Coord, gaps []Coord, gapSize int) int {
	var sum int
	for i := range galaxies {
		for j := i+1; j < len(galaxies); j++ {
			distance := galaxies[j].Distance(galaxies[i])
			for _, gap := range gaps {
				if ((gap.x > 0 &&
					gap.x >= min(galaxies[i].x, galaxies[j].x) &&
					gap.x <= max(galaxies[i].x, galaxies[j].x)) ||
					(gap.y > 0 &&
					gap.y >= min(galaxies[i].y, galaxies[j].y) &&
					gap.y <= max(galaxies[i].y, galaxies[j].y))) {
					distance += gapSize-1
				}
			}
			sum += distance
		}
	}
	return sum
}

func part1(galaxies []Coord, gaps []Coord) int {
	return distances(galaxies, gaps, 2)
}

func part2(galaxies []Coord, gaps []Coord) int {
	return distances(galaxies, gaps, 1000000)
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]

	var gaps []Coord
	var galaxies []Coord
	for y := range lines {
		gap := true
		for x := range lines[y] {
			if lines[y][x] == '#' {
				galaxies = append(galaxies, Coord{x: x, y: y})
				gap = false
			}
		}
		if gap {
			gaps = append(gaps, Coord{x: 0, y: y})
		}
	}
	for x := range lines[0] {
		gap := true
		for y := range lines {
			if lines[y][x] == '#' {
				gap = false
				break
			}
		}
		if gap {
			gaps = append(gaps, Coord{x: x, y: 0})
		}
	}

	fmt.Println(part1(galaxies, gaps))
	fmt.Println(part2(galaxies, gaps))
}
