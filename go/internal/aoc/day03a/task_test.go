package day03a

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle_WithSample(t *testing.T) {
	// Sample data from: https://adventofcode.com/2021/day/3
	input = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010`

	puz := Puzzle{}
	res := puz.Run()

	assert.Equal(t, 198, res)
}

func Benchmark_Run(b *testing.B) {
	p := Puzzle{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}