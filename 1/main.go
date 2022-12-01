package main

import (
	"fmt"
	"sort"

	"github.com/welociraptor/adventofcode-2022/common"
)

func main() {
	totals := make([]int, 1)

	i := 0
	for _, in := range common.MustReadNumbers("1/input.txt") {
		if in != 0 {
			totals[i] += in
		} else {
			totals = append(totals, 0)
			i++
		}
	}
	sort.Ints(totals)

	lastIdx := len(totals) - 1

	fmt.Println("Part 1:", totals[lastIdx])
	fmt.Println("Part 2:", totals[lastIdx]+totals[lastIdx-1]+totals[lastIdx-2])
}
