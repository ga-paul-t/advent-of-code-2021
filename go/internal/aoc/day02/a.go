package day02

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
	return "02a"
}

func (p PuzzleA) Run() int {
	cmds := strings.Split(input, "\n")

	hpos, dep := 0, 0
	for _, cmd := range cmds {
		instr := strings.Split(cmd, " ")
		v, _ := strconv.Atoi(instr[1])

		switch instr[0] {
		case "forward":
			hpos += v
		case "down":
			dep += v
		default:
			dep -= v
		}
	}

	return hpos * dep
}
