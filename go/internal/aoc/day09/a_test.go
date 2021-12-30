package day09

import (
	"testing"

	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc"
	"github.com/stretchr/testify/assert"
)

func TestPuzzleA_Example(t *testing.T) {
	GridSizeX = 10
	GridSizeY = 5

	// Sample data from: https://adventofcode.com/2021/day/9
	input = `2199943210
3987894921
9856789892
8767896789
9899965678`

	puz := PuzzleA{}
	res := puz.Run()

	assert.Equal(t, 15, res)
}

func TestPuzzleA(t *testing.T) {
	GridSizeX = 100
	GridSizeY = 100

	input = aoc.ReadInputFile()

	puz := PuzzleA{}
	res := puz.Run()

	assert.Equal(t, 480, res)
}

func Benchmark_PuzzleA(b *testing.B) {
	p := PuzzleA{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
