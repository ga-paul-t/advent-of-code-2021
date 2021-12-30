package day09

import (
	"testing"
)

// func TestPuzzleB_Example(t *testing.T) {
// 	// Sample data from: https://adventofcode.com/2021/day/9
// 	input = `2199943210
// 3987894921
// 9856789892
// 8767896789
// 9899965678`

// 	puz := PuzzleB{}
// 	res := puz.Run()

// 	assert.Equal(t, 61229, res)
// }

// func TestPuzzleB(t *testing.T) {
// 	input = aoc.ReadInputFile()

// 	puz := PuzzleB{}
// 	res := puz.Run()

// 	assert.Equal(t, 994266, res)
// }

func Benchmark_PuzzleB(b *testing.B) {
	p := PuzzleB{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
