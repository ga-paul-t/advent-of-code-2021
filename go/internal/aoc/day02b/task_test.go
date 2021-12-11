package day02b

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle_WithSample(t *testing.T) {
	// Sample data from: https://adventofcode.com/2021/day/2
	input = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	puz := Puzzle{}
	res := puz.Run()

	assert.Equal(t, 900, res)
}

func Benchmark_Run(b *testing.B) {
	p := Puzzle{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
