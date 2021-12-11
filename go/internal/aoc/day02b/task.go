package day02b

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
	return "02b"
}

func (p Puzzle) Run() int {
	cmds := strings.Split(input, "\n")

	hpos, dep, aim := 0, 0, 0
	for _, cmd := range cmds {
		instr := strings.Split(cmd, " ")
		v, _ := strconv.Atoi(instr[1])

		switch instr[0] {
		case "forward":
			hpos += v
			dep += aim * v
		case "down":
			aim += v
		default:
			aim -= v
		}
	}

	return hpos * dep
}
