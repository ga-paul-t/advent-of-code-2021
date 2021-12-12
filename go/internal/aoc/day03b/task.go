package day03b

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

type Puzzle struct{}

func (p Puzzle) String() string {
	return "02b"
}

func (p Puzzle) Run() int {
	cmds := strings.Split(input, "\n")
	fmt.Println(len(cmds))

	return 230
}
