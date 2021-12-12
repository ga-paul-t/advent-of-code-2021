package day03a

import (
	_ "embed"
	"math"
	"strconv"
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

	bits := len(readings[0])
	gam := make([]byte, bits)
	eps := make([]byte, bits)

	max := int(math.Ceil(float64(len(readings)) / 2.0))

	for b := 0; b < bits; b++ {
		c := 0

		for _, r := range readings {
			if r[b] == '1' {
				c++
			}

			// Break early when we know the most common bit
			if c == max {
				break
			}
		}

		if c == max {
			gam[b] = '1'
			eps[b] = '0'
		} else {
			gam[b] = '0'
			eps[b] = '1'
		}
	}

	gamt, _ := strconv.ParseInt(string(gam[:]), 2, 64)
	epst, _ := strconv.ParseInt(string(eps[:]), 2, 64)

	return int(gamt * epst)
}
