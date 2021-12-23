package day03

import (
	"testing"

	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc"
	"github.com/stretchr/testify/assert"
)

func TestPuzzleA_Example(t *testing.T) {
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

	puz := PuzzleA{}
	res := puz.Run()

	assert.Equal(t, 198, res)
}

func TestPuzzleA(t *testing.T) {
	input = aoc.ReadInputFile()

	puz := PuzzleA{}
	res := puz.Run()

	assert.Equal(t, 3882564, res)
}

func Benchmark_PuzzleA(b *testing.B) {
	p := PuzzleA{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
