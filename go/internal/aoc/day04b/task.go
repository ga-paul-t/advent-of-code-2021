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

type Board struct {
	Numbers map[string]int
	Marks   []byte
	Size    int
}

func NewBoard(in string) *Board {
	// Each line of the board is on a new line
	lines := strings.Split(in, "\n")
	size := len(lines)

	// Populate the bingo board, recording the correct position
	// of each number within the mark slice
	numbers := map[string]int{}
	for i, line := range lines {
		nums := strings.Fields(line)

		for j, n := range nums {
			numbers[n] = i*size + j
		}
	}

	marks := make([]byte, size*size)

	return &Board{
		Numbers: numbers,
		Marks:   marks,
		Size:    size,
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
	row := int(pos / b.Size)
	col := int(pos % b.Size)

	// Check row
	matched := b.Marks[row*b.Size] == '1' &&
		b.Marks[row*b.Size+1] == '1' &&
		b.Marks[row*b.Size+2] == '1' &&
		b.Marks[row*b.Size+3] == '1' &&
		b.Marks[row*b.Size+4] == '1'

	return matched || b.Marks[col] == '1' &&
		b.Marks[col+b.Size] == '1' &&
		b.Marks[col+b.Size*2] == '1' &&
		b.Marks[col+b.Size*3] == '1' &&
		b.Marks[col+b.Size*4] == '1'
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
	boards := make([]*Board, 0, len(brds))
	for _, b := range brds {
		boards = append(boards, NewBoard(b))
	}

	size := boards[0].Size

	// An the grids are square. It isn't possible to win without until after 'n'
	// numbers are drawn
	for i := 0; i < size; i++ {
		for _, b := range boards {
			b.Mark(nums[i])
		}
	}

	// Continue marking, but check for a winning board
	var winNum int
	var winner *Board

	for i := size; i < len(nums) || len(boards) > 0; i++ {
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
	}

	return winNum * winner.SumUnmarked()
}

/*
for i := size; i < len(nums) && len(boards) > 1; i++ {
		boardNo := 0
		for j, b := range boards {
			pos := b.Mark(nums[i])

			if b.Check(pos) {
				// Board has won, so grab winner
				boardNo = j
				winner = b
				break
			}
		}

		if winner != nil {
			winNum, _ = strconv.Atoi(nums[i])

			// Reshape the slice
			copy(boards[boardNo:], boards[boardNo+1:]) // Shift a[i+1:] left one index.
			boards[len(boards)-1] = nil                // Erase last element (write zero value).
			boards = boards[:len(boards)-1]

			fmt.Printf("%#v\n", boards)
		}
	}
*/
