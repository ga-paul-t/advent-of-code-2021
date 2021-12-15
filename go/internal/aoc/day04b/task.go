package day04b

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed input.txt
	input string
)

type Puzzle struct{}

func (p Puzzle) String() string {
	return "04b"
}

func (p Puzzle) Run() int {
	fmt.Println(input)
	return 0
}
