package day01b

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle_WithSample(t *testing.T) {
	// Sample data from: https://adventofcode.com/2021/day/1
	input = `199
200
208
210
200
207
240
269
260
263`

	puz := Puzzle{}
	res := puz.Run()

	assert.Equal(t, 5, res)
}

func Benchmark_Run(b *testing.B) {
	p := Puzzle{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
