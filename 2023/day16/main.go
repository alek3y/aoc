package main

import (
	"os"
	"strings"
	"fmt"
)

type Direction int

const (
	Upward Direction = iota
	Downward
	Rightward
	Leftward
)

type Beam struct {
	x, y int
	facing Direction
}

func (b Beam) Step(tile byte) []Beam {
	beams := []Beam{
		{x: b.x, y: b.y-1, facing: Upward},
		{x: b.x, y: b.y+1, facing: Downward},
		{x: b.x+1, y: b.y, facing: Rightward},
		{x: b.x-1, y: b.y, facing: Leftward},
	}

	switch b.facing {
	case Upward:
		switch tile {
		case '/': return beams[2:3]
		case '\\': return beams[3:]
		case '-': return beams[2:]
		case '|': fallthrough
		default: return beams[:1]
		}
	case Downward:
		switch tile {
		case '/': return beams[3:]
		case '\\': return beams[2:3]
		case '-': return beams[2:]
		case '|': fallthrough
		default: return beams[1:2]
		}
	case Rightward:
		switch tile {
		case '/': return beams[:1]
		case '\\': return beams[1:2]
		case '|': return beams[:2]
		case '-': fallthrough
		default: return beams[2:3]
		}
	case Leftward:
		switch tile {
		case '/': return beams[1:2]
		case '\\': return beams[:1]
		case '|': return beams[:2]
		case '-': fallthrough
		default: return beams[3:]
		}
	}

	return nil
}

func (b Beam) Valid(field [][]byte) bool {
	return b.y >= 0 && b.y < len(field) && b.x >= 0 && b.x < len(field[b.y])
}

func energize(visits []Beam, field [][]byte) int {
	for i := 0; i < len(visits); i++ {
		visit := visits[i]
		for _, beam := range visit.Step(field[visit.y][visit.x]) {
			if beam.Valid(field) {
				var visited bool
				for j := 0; j < i; j++ {
					if visits[j] == beam {
						visited = true
						break
					}
				}
				if !visited {
					visits = append(visits, beam)
				}
			}
		}
	}

	var energized int
	for i := range visits {
		var counted bool
		for j := 0; j < i; j++ {
			if visits[i].x == visits[j].x && visits[i].y == visits[j].y {
				counted = true
			}
		}
		if !counted {
			energized++
		}
	}
	return energized
}

func part1(field [][]byte) int {
	return energize([]Beam{{x: 0, y: 0, facing: Rightward}}, field)
}

func part2(field [][]byte) int {
	var mostEnergetic int
	for i := 0; i < len(field[0]); i++ {
		top := energize([]Beam{{x: i, y: 0, facing: Downward}}, field)
		bottom := energize([]Beam{{x: i, y: len(field)-1, facing: Upward}}, field)
		mostEnergetic = max(mostEnergetic, top, bottom)
	}
	for i := 0; i < len(field); i++ {
		left := energize([]Beam{{x: 0, y: i, facing: Rightward}}, field)
		right := energize([]Beam{{x: len(field[0])-1, y: i, facing: Leftward}}, field)
		mostEnergetic = max(mostEnergetic, left, right)
	}
	return mostEnergetic
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSuffix(string(bytes), "\n"), "\n")

	var field [][]byte
	for _, line := range lines {
		field = append(field, []byte(line))
	}

	fmt.Println(part1(field))
	fmt.Println(part2(field))
}
