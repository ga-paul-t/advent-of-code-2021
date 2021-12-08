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

	inc := 0
	for i := 1; i < len(str); i++ {
		a, _ := strconv.Atoi(str[i])
		b, _ := strconv.Atoi(str[i-1])
		if a > b {
			inc++
		}
	}

	return inc
}
