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

type Tilt int

const (
	North Tilt = 0b00
	East = 0b01
	South = 0b11
	West = 0b10
)

type Field struct {
	platform [][]byte
	transposed bool
}

func (f Field) At(x, y int) *byte {
	if f.transposed {
		return &f.platform[x][y]
	}
	return &f.platform[y][x]
}

func (f *Field) Transpose() {
	f.transposed = !f.transposed
}

func (f Field) Tilt(tilt Tilt) {
	start := int(tilt & 0b1)

	switch tilt {
	case West:
		fallthrough
	case East:
		start *= len(f.platform[0])-1
		f.Transpose()
	case South:
		fallthrough
	case North:
		start *= len(f.platform)-1
	}

	for x := range f.platform[0] {
		var round int
		for offset := range f.platform {
			y := Abs(start - offset)
			switch *f.At(x, y) {
			case '#':
				round = offset+1
			case 'O':
				*f.At(x, y) = '.'
				*f.At(x, Abs(start - round)) = 'O'
				round++
			}
		}
	}

	if f.transposed {
		f.Transpose()
	}
}

func (f Field) Cycle() {
	f.Tilt(North)
	f.Tilt(West)
	f.Tilt(South)
	f.Tilt(East)
}

func (f Field) Load() int {
	var load int
	for x := range f.platform[0] {
		for y := range f.platform {
			if *f.At(x, y) == 'O' {
				load += len(f.platform) - y
			}
		}
	}
	return load
}

func (f Field) Stringify() string {
	var result string
	for _, line := range f.platform {
		result += string(line) + "\n"
	}
	return result
}

func part1(field Field) int {
	field.Tilt(North)
	return field.Load()
}

func part2(field Field) int {
	target := 1000000000

	var loop int
	var cycles []string
	for i := 0; i < target; i++ {
		field.Cycle()
		stringified := field.Stringify()
		cycles = append(cycles, stringified)	// i = len(cycles)-1

		var match bool
		for j := i-1; j >= i/2; j-- {
			if cycles[i] == cycles[j] {
				match = true
				for offset := 0; offset < i - j; offset++ {
					if cycles[j+1 + offset] != cycles[j - (i - j - 1) + offset] {
						match = false
						break
					}
				}

				if match {
					loop = i - j
					break
				}
			}
		}

		if match {
			break
		}
	}

	target -= len(cycles) - 2*loop
	for i := 0; i < target % loop; i++ {
		field.Cycle()
	}
	return field.Load()
}

func main() {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSuffix(string(bytes), "\n"), "\n")

	var field Field
	for _, line := range lines {
		field.platform = append(field.platform, []byte(line))
	}

	fmt.Println(part1(field))
	fmt.Println(part2(field))
}
