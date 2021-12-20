package day02

import (
	"strconv"
	"strings"
)

type PuzzleB struct{}

func (p PuzzleB) String() string {
	return "02b"
}

func (p PuzzleB) Run() int {
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
