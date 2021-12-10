package day01b

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
	return "01b"
}

func (p Puzzle) Run() int {
	depths := strings.Split(input, "\n")

	w := [4]int{}
	w[0], _ = strconv.Atoi(depths[0])
	w[1], _ = strconv.Atoi(depths[1])
	w[2], _ = strconv.Atoi(depths[2])
	w[3], _ = strconv.Atoi(depths[3])

	// Only need to compare the first and last numbers, as each sliding window
	// will share 2 common values:
	//
	// [199, 200, 208, 210] -> A: [199, 299, 208] B: [200, 208, 210]
	// 210 > 199 ++(increased)

	inc := 0
	for i := 4; i < len(depths)-1; i++ {
		if w[3] > w[0] {
			inc++
		}

		// Shift window
		w[0] = w[1]
		w[1] = w[2]
		w[2] = w[3]
		w[3], _ = strconv.Atoi(depths[i+1])
	}

	return inc
}
