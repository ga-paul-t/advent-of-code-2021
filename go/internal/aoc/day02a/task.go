package day02a

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
	return "02a"
}

func (p Puzzle) Run() int {
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
