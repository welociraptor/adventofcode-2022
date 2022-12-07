package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MustReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error reading lines: %w", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func MustReadNumbers(filename string) []int {
	lines := MustReadLines(filename)
	var numbers []int
	for _, in := range lines {
		num, _ := strconv.Atoi(in)
		numbers = append(numbers, num)
	}
	return numbers
}
