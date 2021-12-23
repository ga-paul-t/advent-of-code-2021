package day05

import (
	"testing"

	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc"
	"github.com/stretchr/testify/assert"
)

func TestPuzzleA_Example(t *testing.T) {
	GridSize = 10

	// Sample data from: https://adventofcode.com/2021/day/5
	input = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	puz := PuzzleA{}
	res := puz.Run()

	assert.Equal(t, 5, res)
}

func TestPuzzleA(t *testing.T) {
	GridSize = 1000

	input = aoc.ReadInputFile()

	puz := PuzzleA{}
	res := puz.Run()

	assert.Equal(t, 5169, res)
}

func Benchmark_PuzzleA(b *testing.B) {
	p := PuzzleA{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
