package day05b

import (
	_ "embed"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input string

	// A single dimension of the square grid
	GridSize = 1000
)

type Coords struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func NewCoords(coords []string) Coords {
	c1 := strings.Split(coords[0], ",")
	x1, _ := strconv.Atoi(c1[0])
	y1, _ := strconv.Atoi(c1[1])

	c2 := strings.Split(coords[1], ",")
	x2, _ := strconv.Atoi(c2[0])
	y2, _ := strconv.Atoi(c2[1])

	return Coords{
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}
}

type Puzzle struct{}

func (p Puzzle) String() string {
	return "05b"
}

func (p Puzzle) Run() int {
	grid := make([]int, GridSize*GridSize)

	lines := strings.Split(input, "\n")
	coords := make([]Coords, 0, len(lines))
	for i := range lines {
		coords = append(coords, NewCoords(strings.Split(lines[i], " -> ")))
	}

	overlaps := 0
	plot := func(x, y int) {
		// Map coordinates to position within flat grid
		pos := x + y*GridSize

		if grid[pos] == 1 {
			overlaps++
		}

		grid[pos]++
	}

	for _, c := range coords {
		xinc, yinc := 0, 0

		// Crude, but seems to be the most efficient approach so far!
		if c.X2-c.X1 > 0 {
			xinc = 1
		} else if c.X2-c.X1 < 0 {
			xinc = -1
		}

		if c.Y2-c.Y1 > 0 {
			yinc = 1
		} else if c.Y2-c.Y1 < 0 {
			yinc = -1
		}

		x := c.X1
		y := c.Y1

		for {
			if x == c.X2+xinc && y == c.Y2+yinc {
				break
			}

			plot(x, y)

			x += xinc
			y += yinc
		}
	}

	return overlaps
}
