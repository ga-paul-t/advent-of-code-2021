package day03b

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
	return "03b"
}

func (p Puzzle) Run() int {
	readings := strings.Split(input, "\n")

	bits := len(readings[0])

	oxyRating := oxygenRating(readings, bits)
	co2Rating := co2Rating(readings, bits)

	return int(oxyRating * co2Rating)
}

func oxygenRating(readings []string, bits int) int64 {
	for i := 0; i < bits && len(readings) > 2; i++ {
		max := int(math.Ceil(float64(len(readings)) / 2.0))

		j, k := 0, 0
		var b byte

		for _, r := range readings {
			if r[i] == '1' {
				j++
			} else {
				k++
			}
			// Circuit break as quickly as possible
			if j == max || k == max {
				break
			}
		}

		if j >= k {
			b = '1'
		} else {
			b = '0'
		}

		oxy := []string{}
		for _, r := range readings {
			if r[i] == b {
				oxy = append(oxy, r)
			}
		}

		readings = oxy
	}

	r := readings[1]
	if readings[0][bits-1] == '1' {
		r = readings[0]
	}

	rv, _ := strconv.ParseInt(r, 2, 64)
	return rv
}

func co2Rating(readings []string, bits int) int64 {
	for i := 0; i < bits && len(readings) > 2; i++ {
		max := int(math.Ceil(float64(len(readings)) / 2.0))

		j, k := 0, 0
		var b byte

		for _, r := range readings {
			if r[i] == '1' {
				j++
			} else {
				k++
			}
			// Circuit break as quickly as possible
			if j == max || k == max {
				break
			}
		}

		if j >= k {
			b = '0'
		} else {
			b = '1'
		}

		co2 := []string{}
		for _, r := range readings {
			if r[i] == b {
				co2 = append(co2, r)
			}
		}

		readings = co2
	}

	r := readings[1]
	if readings[0][bits-1] == '0' {
		r = readings[0]
	}

	rv, _ := strconv.ParseInt(r, 2, 64)
	return rv
}