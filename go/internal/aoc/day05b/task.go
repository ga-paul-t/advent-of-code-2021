package day05b

import (
	_ "embed"
)

var (
	//go:embed input.txt
	input string
)

type Puzzle struct{}

func (p Puzzle) String() string {
	return "05b"
}

func (p Puzzle) Run() int {
	return len(input)
}
