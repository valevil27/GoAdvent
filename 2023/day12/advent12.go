package day12

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func Solve() {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Something went wrong")
		return
	}
	fp := filepath.Dir(file)
	inputFile := fmt.Sprintf("%v/input", fp)
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
	raw := parseInput(input)
	totalValid := 0
	for _, d := range raw {
		totalValid += len(validatedPerms(d, totalValid))
	}
	return int64(totalValid)
}

func validatedPerms(d *rawData, totalValid int) []string {
	perms := permutations(d.str)
	validPerms := []string{}
	for _, perm := range perms {
		if validateData(perm, d.data) {
			validPerms = append(validPerms, perm)
		}
	}
	return validPerms
}

type rawData struct {
	str  string
	data []int
}

func parseInput(input string) []*rawData {
	rdata := []*rawData{}
	for _, line := range strings.Split(input, "\n") {
		str, numStr, _ := strings.Cut(line, " ")
		nums := strings.Split(numStr, ",")
		dmgMap := make([]int, 0, len(nums))
		for _, n := range nums {
			i, err := strconv.Atoi(n)
			if err != nil {
				panic(nil)
			}
			dmgMap = append(dmgMap, i)
		}
		rdata = append(rdata, &rawData{str: str, data: dmgMap})
	}
	return rdata
}

func validateData(str string, data []int) bool {
	dmgStr := strings.FieldsFunc(str, func(r rune) bool { return r == '.' })
	if len(dmgStr) != len(data) {
		return false
	}

	for i := range data {
		if data[i] != len(dmgStr[i]) {
			return false
		}
	}
	return true
}

func permutations(s string) []string {
	unknowns := strings.Count(s, "?")
	if unknowns == 0 {
		return []string{s}
	}
	perms := make([]string, 0, 1<<unknowns)
	for i := 0; i < 1<<unknowns; i++ {
		perm := s
		for j := 0; j < unknowns; j++ {
			if i&(1<<j) != 0 {
				perm = strings.Replace(perm, "?", "#", 1)
			} else {
				perm = strings.Replace(perm, "?", ".", 1)
			}
		}
		perms = append(perms, perm)
	}
	return perms
}

func Part2(filepath string) int64 {
	// input := getInput(filepath)
	return 0
}
