package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculatePart1(t *testing.T) {
	lines := []string{
		"A Y",
		"B X",
		"C Z",
	}
	assert.Equal(t, 15, pointsPart1.calculate(lines))
}

func Test_calculatePart2(t *testing.T) {
	lines := []string{
		"A Y",
		"B X",
		"C Z",
	}
	assert.Equal(t, 12, pointsPart2.calculate(lines))
}
