package day04a

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
	// Determine the row and column index of the position within the
	// marks array. Improving speed up of the lookup
	row := int(pos / BoardDim)
	rowPos := row * BoardDim

	col := int(pos % BoardDim)
	colPos := col + BoardDim

	// Check row
	matched := b.Marks[rowPos] == '1' &&
		b.Marks[rowPos+1] == '1' &&
		b.Marks[rowPos+2] == '1' &&
		b.Marks[rowPos+3] == '1' &&
		b.Marks[rowPos+4] == '1'

	return matched || b.Marks[col] == '1' &&
		b.Marks[colPos] == '1' &&
		b.Marks[colPos*2] == '1' &&
		b.Marks[colPos*3] == '1' &&
		b.Marks[colPos*4] == '1'
}

type Puzzle struct{}

func (p Puzzle) String() string {
	return "04a"
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

	// An the grids are square. It isn't possible to win until after 'n'
	// numbers are drawn
	for i := 0; i < BoardDim; i++ {
		for _, b := range boards {
			b.Mark(nums[i])
		}
	}

	// Continue marking, but check for a winning board
	nsize := len(nums)
	num, boardNum := func(start int) (int, int) {
		for ; start < nsize; start++ {
			for i, b := range boards {
				pos := b.Mark(nums[start])

				if b.Check(pos) {
					num, _ := strconv.Atoi(nums[start])
					return num, i
				}
			}
		}
		return 0, 0
	}(BoardDim)

	return num * boards[boardNum].SumUnmarked()
}
