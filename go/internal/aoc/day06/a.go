package day06

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
	return "06a"
}

func (p PuzzleA) Run() int {
	lifecycle := [9]int{}

	initial := strings.Split(input, ",")
	for i := 0; i < len(initial); i++ {
		day, _ := strconv.Atoi(initial[i])
		lifecycle[day]++
	}

	for day := 1; day <= 80; day++ {
		newFish := lifecycle[0]

		for i := 1; i <= 8; i++ {
			lifecycle[i-1] = lifecycle[i]
		}

		lifecycle[8] = newFish
		lifecycle[6] += newFish
	}

	t := 0
	for _, v := range lifecycle {
		t += v
	}

	return t
}
