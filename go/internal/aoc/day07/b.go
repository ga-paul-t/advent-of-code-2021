package day07

import (
	"math"
	"strconv"
	"strings"
)

type PuzzleB struct{}

func (p PuzzleB) String() string {
	return "07b"
}

func (p PuzzleB) Run() int {
	nums := strings.Split(input, ",")

	hpos := make([]int, len(nums))
	for i, n := range nums {
		hpos[i], _ = strconv.Atoi(n)
	}

	// Calculate the mean distance
	sum, mean := 0, 0
	for _, hp := range hpos {
		sum += hp
	}

	// Due to rounding around the mean, sample both bounds and return the smallest total
	mean = sum / len(hpos)

	dists := [2]int{}
	calculateDist := func(m int) int {
		t := 0
		for _, hp := range hpos {
			dist := int(math.Abs(float64(hp - m)))

			// Include incremental fuel increase
			t += dist * (dist + 1) / 2
		}
		return t
	}

	dists[0] = calculateDist(mean)
	dists[1] = calculateDist(mean + 1)

	if dists[0] < dists[1] {
		return dists[0]
	}

	return dists[1]
}
