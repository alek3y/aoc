package main

import (
	"os"
	"strings"
	"fmt"
	"math"
)

type Point[T any] struct {
	x, y T
}

var Pipes = map[byte][]Pipe{
	'|': []Pipe{{x: 0, y: -1}, {x: 0, y: 1}},
	'-': []Pipe{{x: 1, y: 0}, {x: -1, y: 0}},
	'L': []Pipe{{x: 0, y: -1}, {x: 1, y: 0}},
	'J': []Pipe{{x: 0, y: -1}, {x: -1, y: 0}},
	'7': []Pipe{{x: 0, y: 1}, {x: -1, y: 0}},
	'F': []Pipe{{x: 1, y: 0}, {x: 0, y: 1}},
	'S': []Pipe{
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
	},
}

type Pipe Point[int]

func (p Pipe) Bounded(field [][]byte) bool {
	return (p.y >= 0 && p.x >= 0 &&
		p.y < len(field) && p.x < len(field[0]))
}

func (lhs Pipe) Add(rhs Pipe) Pipe {
	return Pipe{x: lhs.x+rhs.x, y: lhs.y+rhs.y}
}

func (lhs Pipe) Connected(rhs Pipe, field [][]byte) bool {
	var from bool
	for _, offset := range Pipes[field[lhs.y][lhs.x]] {
		if lhs.Add(offset) == rhs {
			from = true
			break
		}
	}

	var to bool
	for _, offset := range Pipes[field[rhs.y][rhs.x]] {
		if lhs == rhs.Add(offset) {
			to = true
			break
		}
	}

	return from && to
}

func (p Pipe) Neighbors(field [][]byte) []Pipe {
	var neighbors []Pipe
	for _, offset := range Pipes[field[p.y][p.x]] {
		neighbor := p.Add(offset)
		if neighbor.Bounded(field) && p.Connected(neighbor, field) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func (p Pipe) Path(field [][]byte) ([]Pipe, bool) {
	neighbors := p.Neighbors(field)
	if len(neighbors) == 0 {
		return nil, false
	}

	proceeds := true
	path := []Pipe{p}
	next := neighbors[0]
	for next != p && proceeds {
		proceeds = false
		neighbors = next.Neighbors(field)
		for _, neighbor := range neighbors {
			if neighbor != path[len(path)-1] {
				proceeds = true
				path = append(path, next)
				next = neighbor
				break
			}
		}
	}
	return path, proceeds
}

func part1(start Pipe, field [][]byte) int {
	path, _ := start.Path(field)
	return len(path)/2
}

type Ground Point[float64]

func (g Ground) Floor() Pipe {
	return Pipe{x: int(math.Floor(g.x)), y: int(math.Floor(g.y))}
}

func (g Ground) Ceil() Pipe {
	return Pipe{x: int(math.Ceil(g.x)), y: int(math.Ceil(g.y))}
}

func (g Ground) Edge(field [][]byte) bool {
	return (g.y == 0 || g.y == float64(len(field)-1) ||
		g.x == 0 || g.x == float64(len(field[0])-1))
}

func (g Ground) Ground(loop []Pipe, field [][]byte) bool {
	floor, ceil := g.Floor(), g.Ceil()
	if floor == ceil {
		for _, pipe := range loop {
			if pipe == ceil {
				return false
			}
		}
		return true
	} else if ceil.y < len(field) && ceil.x < len(field[ceil.y]) {
		return !floor.Connected(ceil, field)
	}
	return false
}

func (g Ground) Neighbors(loop []Pipe, field [][]byte) []Ground {
	var neighbors []Ground
	for y := max(g.y-0.5, 0); y <= min(g.y+0.5, float64(len(field))-1); y += 0.5 {
		for x := max(g.x-0.5, 0); x <= min(g.x+0.5, float64(len(field[0]))-1); x += 0.5 {
			neighbor := Ground{x: x, y: y}
			if neighbor != g && neighbor.Ground(loop, field) {
				neighbors = append(neighbors, neighbor)
			}
		}
	}
	return neighbors
}

func part2(start Pipe, field [][]byte) int {
	var sum int
	loop, _ := start.Path(field)
	for y := range field {
		for x := range field[0] {
			if field[y][x] == '.' {
				var grounds int

				var visit int
				visits := []Ground{{x: float64(x), y: float64(y)}}
				for visit < len(visits) && !visits[visit].Edge(field) {
					ground := visits[visit]
					visit++

					if ground.Ceil() == ground.Floor() {
						grounds++
					}

					for _, neighbor := range ground.Neighbors(loop, field) {
						var added bool
						for _, enqueued := range visits {
							if neighbor == enqueued {
								added = true
								break
							}
						}
						if !added {
							visits = append(visits, neighbor)
						}
					}
				}

				if visit == len(visits) {
					for _, ground := range visits {
						if ceil := ground.Ceil(); ceil == ground.Floor() {
							field[ceil.y][ceil.x] = 'I'
						}
					}
					sum += grounds
				}
			}
		}
	}
	return sum
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(bytes), "\n")
	lines = lines[:len(lines)-1]

	var field [][]byte
	var start Pipe
	for y := range lines {
		x := strings.Index(lines[y], "S")
		if x >= 0 {
			start = Pipe{x: x, y: y}
		}
		field = append(field, []byte(lines[y]))
	}

	fmt.Println(part1(start, field))
	fmt.Println(part2(start, field))
}
