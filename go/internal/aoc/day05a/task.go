package day05a

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

	// Simplify calculations later by ordering the points to ensure all steps are positive.
	// This does break diagonal coordinates, but these are ignored
	c := Coords{
		X1: x1,
		Y1: y2,
		X2: x2,
		Y2: y2,
	}

	if x1 == x2 && y1 > y2 {
		c.Y1 = y2
		c.Y2 = y1
	} else if y1 == y2 && x1 > x2 {
		c.X1 = x2
		c.X2 = x1
	}

	return c
}

type Puzzle struct{}

func (p Puzzle) String() string {
	return "05a"
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
			return
		}

		grid[pos] = 1
	}

	for _, c := range coords {
		// Determine which axis to plot along
		if c.X1 == c.X2 {
			for y := c.Y1; y <= c.Y2; y++ {
				plot(c.X1, y)
			}
		} else if c.Y1 == c.Y2 {
			for x := c.X1; x <= c.X2; x++ {
				plot(x, c.Y1)
			}
		}
	}

	return overlaps
}
