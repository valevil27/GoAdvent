package day09

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	inputFile := "./2023/day09/input"
	fmt.Printf("inputFile: %v\n", inputFile)
	fmt.Printf("Part1(inputFile): %v\n", Part1(inputFile))
	fmt.Printf("Part2(inputFile): %v\n", Part2(inputFile))
}

func getInput(filepath string) string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func Part1(filepath string) uint64 {
	input := getInput(filepath)
	histories := parseInput(input)
	// For history get the result
	sum := 0
	for _, h := range histories {
		fmt.Printf("h: %v\n", h)
		fs := getFullStory(h)
		fmt.Printf("fs: %v\n", fs)
		ll := getForecast(fs)
		fmt.Printf("ll: %v\n", ll)
		sum += ll
	}
	// sum the results
	return uint64(sum)
}

func parseInput(input string) [][]int {
	histories := [][]int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		numbers := []int{}
		for _, n := range strings.Split(line, " ") {
			nint, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, nint)
		}
		histories = append(histories, numbers)
	}
	return histories
}

func getIncreasingPattern(h []int) []int {
	if len(h) < 2 {
		panic("Too few inputs!")
	}
	increments := []int{}
	for i := 1; i < len(h); i++ {
		increments = append(increments, h[i]-h[i-1])
	}
	return increments
}

func isLastIncrement(h []int) bool {
	for _, v := range h {
		if v != 0 {
			return false
		}
	}
	return true
}

func getFullStory(h []int) []int {
	full := []int{}
	current := h
	for {
		full = append(full, current[len(current)-1])
		current = getIncreasingPattern(current)
		if isLastIncrement(current) {
			return full
		}
	}
}

func getForecast(full []int) int {
	sum := 0
	for _, i := range full {
		sum += i
	}
	return sum
}

func Part2(filepath string) uint64 {
	// input := getInput(filepath)
	return uint64(0)
}
