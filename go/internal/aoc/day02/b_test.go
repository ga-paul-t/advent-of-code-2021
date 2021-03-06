package day02

import (
	"testing"

	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc"
	"github.com/stretchr/testify/assert"
)

func TestPuzzleB_Example(t *testing.T) {
	// Sample data from: https://adventofcode.com/2021/day/2
	input = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	puz := PuzzleB{}
	res := puz.Run()

	assert.Equal(t, 900, res)
}

func TestPuzzleB(t *testing.T) {
	input = aoc.ReadInputFile()

	puz := PuzzleB{}
	res := puz.Run()

	assert.Equal(t, 1975421260, res)
}

func Benchmark_PuzzleB(b *testing.B) {
	p := PuzzleB{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
