package day04b

import (
	_ "embed"
	"strconv"
	"strings"
)

var (
	//go:embed input.txt
	input string
)

const (
	BoardDim  = 5
	BoardSize = 25
)

type Board struct {
	Numbers map[string]int
	Marks   [BoardSize]byte
}

func NewBoard(in string) *Board {
	nums := strings.Fields(in)

	// Populate the bingo board, recording the correct position
	// of each number within the mark slice
	numbers := map[string]int{}
	for i := range nums {
		numbers[nums[i]] = i
	}

	return &Board{
		Numbers: numbers,
	}
}

func (b *Board) Mark(num string) int {
	i, ok := b.Numbers[num]
	if !ok {
		return 0
	}
	b.Marks[i] = '1'

	// Remove marked number from the board
	delete(b.Numbers, num)
	return i
}

func (b *Board) SumUnmarked() int {
	sum := 0
	for k := range b.Numbers {
		n, _ := strconv.Atoi(k)
		sum += n
	}
	return sum
}

func (b *Board) Check(pos int) bool {
	// Determine which row the position is within the board
	row := int(pos / BoardDim)
	col := int(pos % BoardDim)

	// Check row
	matched := b.Marks[row*BoardDim] == '1' &&
		b.Marks[row*BoardDim+1] == '1' &&
		b.Marks[row*BoardDim+2] == '1' &&
		b.Marks[row*BoardDim+3] == '1' &&
		b.Marks[row*BoardDim+4] == '1'

	return matched || b.Marks[col] == '1' &&
		b.Marks[col+BoardDim] == '1' &&
		b.Marks[col+BoardDim*2] == '1' &&
		b.Marks[col+BoardDim*3] == '1' &&
		b.Marks[col+BoardDim*4] == '1'
}

type Puzzle struct{}

func (p Puzzle) String() string {
	return "04b"
}

func (p Puzzle) Run() int {
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
