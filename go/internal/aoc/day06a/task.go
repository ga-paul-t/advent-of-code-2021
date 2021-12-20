package day06a

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
	return "06a"
}

func (p Puzzle) Run() int {
	lifecycle := map[int]int{
		0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0,
	}

	initial := strings.Split(input, ",")
	for i := 0; i < len(initial); i++ {
		day, _ := strconv.Atoi(initial[i])
		lifecycle[day]++
	}

	for day := 1; day <= 80; day++ {
		newFish := 0
		for i := 0; i <= 8; i++ {
			if i == 0 {
				newFish = lifecycle[i]
			} else {
				lifecycle[i-1] = lifecycle[i]
			}
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
