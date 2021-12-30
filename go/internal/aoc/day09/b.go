package day09

import (
	"sort"
	"strings"
)

type PuzzleB struct{}

func (p PuzzleB) String() string {
	return "09b"
}

func (p PuzzleB) Run() int {
	grid := make([][]int, GridSizeY)
	for i := range grid {
		grid[i] = make([]int, GridSizeX)
	}

	lines := strings.Split(input, "\n")
	for i := range lines {
		for j, c := range lines[i] {
			grid[i][j] = int(c - '0')
		}
	}

	// Just wanted to see if this possible in Go, using recursion on an anonymous function
	var searchBasin func(x, y int) int
	searchBasin = func(x, y int) int {
		c := 0

		// Recursively search all neighbours, marking each visited point with a 9
		// to ensure it is not traversed again
		grid[y][x] = 9

		// Top
		if y >= 1 && grid[y-1][x] < 9 {
			c += searchBasin(x, y-1) + 1
		}

		// Left
		if x >= 1 && grid[y][x-1] < 9 {
			c += searchBasin(x-1, y) + 1
		}

		// Right
		if x < GridSizeX-1 && grid[y][x+1] < 9 {
			c += searchBasin(x+1, y) + 1
		}

		// Down
		if y < GridSizeY-1 && grid[y+1][x] < 9 {
			c += searchBasin(x, y+1) + 1
		}

		return c
	}

	basins := []int{}
	for y := range grid {
		for x := range grid[y] {
			// A basin is identified by the first non 9 number
			if grid[y][x] < 9 {
				b := searchBasin(x, y) + 1
				basins = append(basins, b)
			}
		}
	}

	sort.Ints(basins)

	t := basins[len(basins)-1]
	for _, v := range basins[len(basins)-3 : len(basins)-1] {
		t *= v
	}

	return t
}
