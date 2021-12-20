package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzleA_WithSample(t *testing.T) {
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

	puz := PuzzleA{}
	res := puz.Run()

	assert.Equal(t, 7, res)
}

func Benchmark_PuzzleA(b *testing.B) {
	p := PuzzleA{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
