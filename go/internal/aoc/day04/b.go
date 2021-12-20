package day04

import (
	"strconv"
	"strings"
)

type PuzzleB struct{}

func (p PuzzleB) String() string {
	return "04b"
}

func (p PuzzleB) Run() int {
	index := strings.Index(input, "\n\n")

	// Parse the input, splitting the numbers from the boards
	nums := strings.Split(input[:index], ",")
	brds := strings.Split(input[index+2:], "\n\n")

	// Parse boards from input and convert into required structure
	boards := make([]*Board, len(brds))
	for i, b := range brds {
		boards[i] = NewBoard(b)
	}

	// An the grids are square. It isn't possible to win without until after 'n'
	// numbers are drawn
	for i := 0; i < BoardDim; i++ {
		for _, b := range boards {
			b.Mark(nums[i])
		}
	}

	// Continue marking, but check for a winning board
	var winNum int
	var winner *Board

	nsize := len(nums)
	for i := BoardDim; i < nsize; i++ {
		for n := 0; n < len(boards); n++ {
			board := boards[n]
			pos := board.Mark(nums[i])

			if board.Check(pos) {
				winNum, _ = strconv.Atoi(nums[i])
				winner = board

				// Remove from the board from the slice
				boards = append(boards[:n], boards[n+1:]...)
				n--
			}
		}

		if len(boards) == 0 {
			break
		}
	}

	return winNum * winner.SumUnmarked()
}
