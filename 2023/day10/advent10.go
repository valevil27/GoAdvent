package day10

import (
	"fmt"
	"os"
	"strings"
)

func Solve() {
	inputFile := "./2023/day10/input"
	fmt.Printf("inputFile: %v\n", inputFile)
	fmt.Printf("Part1(input): %v\n", Part1(inputFile))
	fmt.Printf("Part2(input): %v\n", Part2(inputFile))
}

func getInput(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func Part1(filepath string) int64 {
	input := getInput(filepath)
	tiles, starting := parseInput(input)
	if starting == [2]int{0, 0} {
		panic("Help, starting point is nil!")
	}
	var solution int
	for _, d := range []Direction{North, East, South, West} {
		isPath, pathLength := getPath(tiles, starting, d)
		if isPath {
			solution = pathLength/2 + 1
			break
		}
	}

	return int64(solution)
}

func getPath(tiles [][]*Pipe, starting [2]int, direction Direction) (bool, int) {
	x, y := starting[0], starting[1]
	pathLength := 0
	var nextPipe *Pipe
	for {
		switch direction {
		case North:
			y = y - 1
		case East:
			x = x + 1
		case South:
			y = y + 1
		case West:
			x = x - 1
		case None:
			return false, 0
		}
		if x < 0 || x >= len(tiles[0]) || y < 0 || y >= len(tiles) {
			return false, 0
		}
		nextPipe = tiles[y][x]
		if nextPipe.start {
			return true, pathLength
		}
		if (direction+2)%4 == nextPipe.from {
			pathLength++
			direction = nextPipe.to
		} else if (direction+2)%4 == nextPipe.to {
			pathLength++
			direction = nextPipe.from
		} else {
			return false, 0
		}
	}
}

type Direction byte

const (
	North Direction = iota
	East
	South
	West
	None
)

type Pipe struct {
	// tile rune
	from  Direction
	to    Direction
	start bool
}

var mapping map[rune]*Pipe = map[rune]*Pipe{
	'.': {from: None, to: None},
	'S': {from: None, to: None, start: true},
	'|': {from: North, to: South},
	'-': {from: East, to: West},
	'L': {from: North, to: East},
	'J': {from: North, to: West},
	'7': {from: South, to: West},
	'F': {from: South, to: East},
}

func parseInput(input string) ([][]*Pipe, [2]int) {
	tiles := [][]*Pipe{}
	var starting [2]int = [2]int{0, 0}
	for y, l := range strings.Split(input, "\n") {
		tile_line := []*Pipe{}
		for x, r := range l {
			tile_line = append(tile_line, mapping[r])
			if r == 'S' {
				starting = [2]int{x, y}
			}
		}
		tiles = append(tiles, tile_line)
	}
	return tiles, starting
}

func Part2(filepath string) int64 {
	// input := getInput(filepath)
	return 0
}
