package day07

import (
	_ "embed"
	"math"
	"sort"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "07a"
}

func (p PuzzleA) Run() int {
	nums := strings.Split(input, ",")

	hpos := make([]int, len(nums))
	for i, n := range nums {
		hpos[i], _ = strconv.Atoi(n)
	}

	// Determine the median from the set of horizontal positions
	sort.Ints(hpos)
	mid := len(hpos) / 2
	med := hpos[mid]

	t := 0
	for _, hp := range hpos {
		t += int(math.Abs(float64(hp - med)))
	}

	return t
}
