package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzleB_WithSample(t *testing.T) {
	// Sample data from: https://adventofcode.com/2021/day/6
	input = "3,4,3,1,2"

	puz := PuzzleB{}
	res := puz.Run()

	assert.Equal(t, 26984457539, res)
}

func Benchmark_PuzzleB(b *testing.B) {
	p := PuzzleB{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
