package main

import (
	"fmt"

	"github.com/welociraptor/adventofcode-2022/common"
)

func main() {
	lines := common.MustReadLines("2/input.txt")

	fmt.Println("part 1:", pointsPart1.calculate(lines))
	fmt.Println("part 2:", pointsPart2.calculate(lines))
}

type Scoring map[string]int

func (s Scoring) calculate(args []string) int {
	total := 0
	for _, line := range args {
		total += s[line]
	}
	return total
}

var pointsPart1 = Scoring{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
}

var pointsPart2 = Scoring{
	"A X": 3,
	"A Y": 4,
	"A Z": 8,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 2,
	"C Y": 6,
	"C Z": 7,
}
