package day08

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	math_utils "github.com/valevil27/adventofgo/utils/math"
)

func Solve() {
	inputFile := "./2023/day08/input"
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
	dirs, maps, _ := strings.Cut(input, "\n\n")
	nodes := parseNodes(maps)
	current_node := nodes["AAA"]
	steps := 0
	for {
		for _, d := range dirs {
			steps++
			if d == 'L' {
				current_node = current_node.left
			} else {
				current_node = current_node.right
			}
			if current_node.tag == "ZZZ" {
				break
			}
		}
		if current_node.tag == "ZZZ" {
			break
		}
	}
	return uint64(steps)
}

func Part2(filepath string) uint64 {
	input := getInput(filepath)
	dirs, maps, _ := strings.Cut(input, "\n\n")
	nodes := parseNodes(maps)
	// Find starting nodes
	startingTags := getStartingTags(nodes)

	// Solve for each starting node
	allSteps := []int{}
	for _, tag := range startingTags {
		currentNode := nodes[tag]
		steps := 0
		for {
			for _, d := range dirs {
				steps++
				if d == 'L' {
					currentNode = currentNode.left
				} else {
					currentNode = currentNode.right
				}
				if isEndingTag(currentNode.tag) {
					break
				}
			}
			if isEndingTag(currentNode.tag) {
				break
			}
		}
		allSteps = append(allSteps, steps)
	}
	// Find mcm for steps
	fmt.Printf("allSteps: %v\n", allSteps)
	return uint64(math_utils.Lcm(allSteps))
}

type Node struct {
	tag   string
	left  *Node
	right *Node
}

func parseNodes(maps string) map[string]*Node {
	nodes := make(map[string]*Node)
	for _, line := range strings.Split(maps, "\n") {
		current := strings.Split(line, " = ")[0]
		nodes[current] = &Node{tag: current}
	}
	regpat := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	for _, line := range strings.Split(maps, "\n") {
		matches := regpat.FindStringSubmatch(line)
		nodes[matches[1]].left = nodes[matches[2]]
		nodes[matches[1]].right = nodes[matches[3]]
	}
	return nodes
}

func getStartingTags(nodes map[string]*Node) []string {
	startingTags := []string{}
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			startingTags = append(startingTags, k)
		}
	}
	return startingTags
}

func isEndingTag(tag string) bool {
	return strings.HasSuffix(tag, "Z")
}
