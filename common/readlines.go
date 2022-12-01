package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func MustReadNumbers(filename string) []int {
	lines, err := ReadLines(filename)
	if err != nil {
		fmt.Println("error reading lines: %w", err)
		os.Exit(1)
	}
	var numbers []int
	for _, in := range lines {
		num, _ := strconv.Atoi(in)
		numbers = append(numbers, num)
	}
	return numbers
}
