package day06

import (
	"testing"

	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc"
	"github.com/stretchr/testify/assert"
)

func TestPuzzleB_Example(t *testing.T) {
	// Sample data from: https://adventofcode.com/2021/day/6
	input = "3,4,3,1,2"

	puz := PuzzleB{}
	res := puz.Run()

	assert.Equal(t, 26984457539, res)
}

func TestPuzzleB(t *testing.T) {
	input = aoc.ReadInputFile()

	puz := PuzzleB{}
	res := puz.Run()

	assert.Equal(t, 1590327954513, res)
}

func Benchmark_PuzzleB(b *testing.B) {
	p := PuzzleB{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
