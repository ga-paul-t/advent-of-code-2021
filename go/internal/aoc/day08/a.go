package day08

import (
	_ "embed"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "08a"
}

func (p PuzzleA) Run() int {
	lines := strings.Split(input, "\n")

	sigs := make([]string, 0, len(lines)*4)
	for _, line := range lines {
		sigs = append(sigs, strings.Fields(strings.Split(line, " | ")[1])...)
	}

	// Count unique signal patterns based on length of string
	c := 0
	for _, sig := range sigs {
		l := len(sig)
		if l == 2 || l == 3 || l == 4 || l == 7 {
			c++
		}
	}

	return c
}
