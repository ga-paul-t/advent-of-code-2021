package day06a

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle_WithSample(t *testing.T) {
	// Sample data from: https://adventofcode.com/2021/day/6
	input = "3,4,3,1,2"

	puz := Puzzle{}
	res := puz.Run()

	assert.Equal(t, 5934, res)
}

func Benchmark_Run(b *testing.B) {
	p := Puzzle{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
