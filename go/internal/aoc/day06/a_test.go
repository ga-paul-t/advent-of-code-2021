package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzleA_WithSample(t *testing.T) {
	// Sample data from: https://adventofcode.com/2021/day/6
	input = "3,4,3,1,2"

	puz := PuzzleA{}
	res := puz.Run()

	assert.Equal(t, 5934, res)
}

func Benchmark_PuzzleA(b *testing.B) {
	p := PuzzleA{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
