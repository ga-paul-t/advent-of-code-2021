package day01a

import (
	_ "embed"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

type Puzzle struct{}

func (p Puzzle) String() string {
	return "01a"
}

func (p Puzzle) Run() int {
	str := strings.Split(input, "\n")

	left, _ := strconv.Atoi(str[0])
	right, inc := 0, 0
	for i := 1; i < len(str); i++ {
		right, _ = strconv.Atoi(str[i])
		if right > left {
			inc++
		}
		left = right
	}

	return inc
}
