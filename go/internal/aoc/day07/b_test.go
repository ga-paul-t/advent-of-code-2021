package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzleB_WithSample(t *testing.T) {
	// Sample data from: https://adventofcode.com/2021/day/7
	input = "16,1,2,0,4,2,7,1,2,14"

	puz := PuzzleB{}
	res := puz.Run()

	assert.Equal(t, 168, res)
}

func Benchmark_PuzzleB(b *testing.B) {
	p := PuzzleB{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
