package day07

import (
	"testing"

	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc"
	"github.com/stretchr/testify/assert"
)

func TestPuzzleA_Example(t *testing.T) {
	// Sample data from: https://adventofcode.com/2021/day/7
	input = "16,1,2,0,4,2,7,1,2,14"

	puz := PuzzleA{}
	res := puz.Run()

	assert.Equal(t, 37, res)
}

func TestPuzzleA(t *testing.T) {
	input = aoc.ReadInputFile()

	puz := PuzzleA{}
	res := puz.Run()

	assert.Equal(t, 329389, res)
}

func Benchmark_PuzzleA(b *testing.B) {
	p := PuzzleA{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
