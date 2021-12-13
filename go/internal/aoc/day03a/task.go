package day03a

import (
	_ "embed"
	"math"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

type Puzzle struct{}

func (p Puzzle) String() string {
	return "03a"
}

func (p Puzzle) Run() int {
	readings := strings.Split(input, "\n")

	// Determine the bias for setting a bit to 1
	max := int(math.Ceil(float64(len(readings)) / 2.0))

	bits := len(readings[0])
	gam, eps := 0, 0

	for b := 0; b < bits; b++ {
		c := 0

		for _, r := range readings {
			if r[b] == '1' {
				c++
			}
		}

		if c >= max {
			gam += 1 << (bits - (b + 1))
		}
	}

	// Invert gamma
	eps = ^gam & ((1 << bits) - 1)

	return gam * eps
}
