package day07

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve07() {
	fmt.Println("Solving exercise 7 part 1")
	hands := parseInput("2023/07/test-input")
	for _, h := range hands {
		fmt.Print(h.cards, h.type_, h.joker)
	}
	fmt.Println()
	mapping := map[byte]byte{'T': 'A', 'J': 'B', 'Q': 'C', 'K': 'D', 'A': 'E'}
	sortHands(hands, mapping)
	for _, h := range hands {
		fmt.Print(h.cards)
	}
	fmt.Println()
	var sum uint64 = 0
	for i, h := range hands {
		sum += h.bid * uint64(i+1)
	}
	fmt.Println(sum)

	fmt.Println("Solving exercise 7 part 2")
	hands = parseInput("2023/07/input")
	mapping = map[byte]byte{'T': 'A', 'J': '.', 'Q': 'C', 'K': 'D', 'A': 'E'}
	sortJokerHands(hands, mapping)
	for _, h := range hands {
		fmt.Print(h.cards)
	}
	fmt.Println()
	sum = 0
	for i, h := range hands {
		sum += h.bid * uint64(i+1)
	}
	fmt.Println(sum)
}

type HandType byte

const (
	High HandType = iota
	One_pair
	Two_pair
	Three_kind
	Full_house
	Four_kind
	Five_kind
)

type Hand struct {
	bid   uint64
	type_ HandType
	joker HandType
	cards []byte
}

func parseInput(filename string) []Hand {
	input, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()
	fscanner := bufio.NewScanner(input)
	hands := make([]Hand, 0)
	// 32T3K 765
	for fscanner.Scan() {
		line := strings.Split(fscanner.Text(), " ")
		bid, _ := strconv.Atoi(line[1])
		cards := []byte(line[0])
		type_ := getHandType(cards)
		joker := getJokerType(cards, type_)
		hand := Hand{bid: uint64(bid), cards: cards, type_: type_, joker: joker}
		hands = append(hands, hand)
	}
	return hands
}

func getHandType(cards []byte) HandType {
	counts := countReps(cards)
	switch len(counts) {
	case 1:
		return Five_kind
	case 2:
		if slices.Contains(counts, 4) {
			return Four_kind
		}
		return Full_house
	case 3:
		if slices.Contains(counts, 3) {
			return Three_kind
		}
		return Two_pair
	case 4:
		return One_pair
	default:
		return High
	}
}

func countReps(arr []byte) []int {
	values := map[byte]int{}
	for _, v := range arr {
		_, ok := values[v]
		if ok {
			values[v]++
		} else {
			values[v] = 1
		}
	}
	counts := make([]int, 0)
	for _, v := range values {
		counts = append(counts, v)
	}
	return counts
}

func sortHands(hands []Hand, mapping map[byte]byte) {
	for _, h := range hands {
		for i, c := range h.cards {
			mappedChar, ok := mapping[c]
			if ok {
				h.cards[i] = mappedChar
			}
		}
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.type_ > b.type_ {
			return 1
		}
		if a.type_ < b.type_ {
			return -1
		}
		for i := range a.cards {
			if a.cards[i] > b.cards[i] {
				return 1
			}
			if a.cards[i] < b.cards[i] {
				return -1
			}
		}
		return 0
	})
}

func getJokerType(cards []byte, type_ HandType) HandType {
	js := 0
	for _, c := range cards {
		if c == 'J' {
			js++
		}
	}
	if js == 0 {
		return type_
	}
	switch type_ {
	case Five_kind, Four_kind, Full_house:
		return Five_kind
	case Three_kind:
		return Four_kind
	case Two_pair:
		if js == 2 {
			return Four_kind
		}
		return Full_house
	case One_pair:
		return Three_kind
	default:
		return One_pair
	}
}

func sortJokerHands(hands []Hand, mapping map[byte]byte) {
	for _, h := range hands {
		for i, c := range h.cards {
			mappedChar, ok := mapping[c]
			if ok {
				h.cards[i] = mappedChar
			}
		}
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.joker > b.joker {
			return 1
		}
		if a.joker < b.joker {
			return -1
		}
		for i := range a.cards {
			if a.cards[i] > b.cards[i] {
				return 1
			}
			if a.cards[i] < b.cards[i] {
				return -1
			}
		}
		return 0
	})
}
