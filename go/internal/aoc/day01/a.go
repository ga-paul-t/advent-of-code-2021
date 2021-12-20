package day01

import (
	_ "embed"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "01a"
}

func (p PuzzleA) Run() int {
	depths := strings.Split(input, "\n")

	a, _ := strconv.Atoi(depths[0])
	b, inc := 0, 0
	for i := 1; i < len(depths); i++ {
		b, _ = strconv.Atoi(depths[i])
		if b > a {
			inc++
		}
		a = b
	}

	return inc
}
