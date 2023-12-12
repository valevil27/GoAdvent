package day11

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func Solve() {
	inputFile := "./2023/day11/input"
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
	// 2 - For all rows and cols
	// 2a - Find Galaxies -> Array
	universe := []*Galaxy{}
	nGalaxy := 1
	splitInput := strings.Split(input, "\n")
	for y, row := range splitInput {
		for x, ch := range row {
			if ch == '#' {
				universe = append(universe, &Galaxy{n: nGalaxy, x: x, y: y})
				nGalaxy++
			}
		}
	}
	// 1 - Perform expansions, order doesn't really matter
	// 1a - rows
	y := 0
	for _, row := range splitInput {
		if !strings.Contains(row, "#") {
			for _, g := range universe {
				if g.y > y {
					g.y++
				}
			}
			y++
		}
		y++
	}
	// 1b - cols
	x := 0
	for c := range splitInput[0] {
		col := ""
		for _, row := range splitInput {
			col += string(row[c])
		}
		if !strings.Contains(col, "#") {
			for _, g := range universe {
				if g.x > x {
					g.x++
				}
			}
			x++
		}
		x++
	}
	sumDiffs := 0.
	for i, g := range universe {
		for _, g2 := range universe[i+1:] {
			currDiff := math.Abs(float64(g.x)-float64(g2.x)) + math.Abs(float64(g.y)-float64(g2.y))
			sumDiffs += currDiff
		}
	}

	// 3 - For each galaxy, find its distance to others

	return int64(sumDiffs)
}

type Galaxy struct {
	n    int
	x, y int
}

func Part2(filepath string) int64 {
	// input := getInput(filepath)
	return 0
}
