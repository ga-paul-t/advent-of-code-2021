package day09

import (
	_ "embed"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input string

	// A single dimension of the square grid
	GridSizeX = 100
	GridSizeY = 100
)

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "09a"
}

func (p PuzzleA) Run() int {
	grid := make([][]byte, GridSizeY)
	for i := range grid {
		grid[i] = make([]byte, 0, GridSizeX)
	}

	lines := strings.Split(input, "\n")
	for i := range lines {
		grid[i] = append(grid[i], []byte(lines[i])...)
	}

	t := 0

	// Loop through Grid identifying the neighbours of each point
	var n byte
	for y := range grid {
		for x := range grid[y] {
			n = 0x39

			// Up
			if y >= 1 && grid[y-1][x] < n {
				n = grid[y-1][x]
			}

			// Left
			if x >= 1 && grid[y][x-1] < n {
				n = grid[y][x-1]
			}

			// Right
			if x < GridSizeX-1 && grid[y][x+1] < n {
				n = grid[y][x+1]
			}

			// Down
			if y < GridSizeY-1 && grid[y+1][x] < n {
				n = grid[y+1][x]
			}

			// Is point smaller than all its neighbours
			if grid[y][x] < n {
				v, _ := strconv.Atoi(string(grid[y][x]))
				t += v + 1
			}
		}
	}

	return t
}
