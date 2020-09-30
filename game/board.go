package game

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const (
	_rows, _cols         = 4, 4      // game board size
	_clearScreenSequence = "\033[2J" // used to clear the screen
	_probabilitySpace    = 100
	_probabilityOfTwo    = 80
)

type Board interface {
	Display()
	AddElement()
	TakeInput()
	IsOver()
	CountScore() (int, int)
}

type board struct {
	matrix         [][]int
	over           bool
	newRow, newCol int
}

func (b *board) CountScore() (int, int) {
	total, max := 0, 0
	matrix := b.matrix
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			total += matrix[i][j]
			max = maxInts(max, matrix[i][j])
		}
	}
	return max, total
}

func maxInts(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (b *board) IsOver() bool {
	emptyCount := 0
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.matrix[i][j] == 0 {
				emptyCount++
			}
		}
	}
	return emptyCount == 0 || b.over
}

func (b *board) TakeInput() {
	/* reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') */
	dir, err := getCharKeystroke()
	if err != nil {
		if errors.Is(err, errors.New("GameOverError")) {
			b.over = true
			return
		} else {
			panic(err)
		}
	}
	//fmt.Printf("the dir is: %v \n", dir)
	if dir == NO_DIR {
		b.TakeInput()
	}
	switch dir {
	case LEFT:
		b.moveLeft()
	case RIGHT:
		b.moveRight()
	case UP:
		b.moveUp()
	case DOWN:
		b.moveDown()
	}
}

func (b *board) AddElement() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// generate new element
	val := r.Int() % _probabilitySpace
	if val < _probabilityOfTwo {
		val = 2
	} else {
		val = 4
	}

	// count empty postions
	emptyCount := 0
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.matrix[i][j] == 0 {
				emptyCount++
			}
		}
	}

	// generate the element position to be filled
	elementCount := r.Int()%emptyCount + 1
	index := 0

	// fill
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.matrix[i][j] == 0 {
				index++
				if index == elementCount {
					b.newRow, b.newCol = i, j
					b.matrix[i][j] = val
					return
				}
			}
		}
	}
}

/* Display board as follows
------------------------------------------------
  2048     |    16     |  1024     |    16
------------------------------------------------
   128     |    16     |    16     |   128
------------------------------------------------
    32     |   512     |   256     |    64
------------------------------------------------
   256     |     4     |   256     |    32
------------------------------------------------
*/
func (b *board) Display() {
	fmt.Println(_clearScreenSequence)
	d := color.New(color.FgBlue, color.Bold)
	printHorizontal()
	for i := 0; i < len(b.matrix); i++ {
		for j := 0; j < len(b.matrix[0]); j++ {
			if b.matrix[i][j] == 0 {
				fmt.Printf("%6s", "")
			} else {
				if i == b.newRow && j == b.newCol {
					d.Printf("%6d", b.matrix[i][j])
				} else {
					fmt.Printf("%6d", b.matrix[i][j])
				}
			}
			fmt.Printf("%5s", "")
			if j != len(b.matrix[0])-1 {
				printVertical()
			}
		}
		fmt.Printf("%5s", "")
		fmt.Println()
		printHorizontal()
	}
}

// printHorizontal prints a grid row
func printHorizontal() {
	for i := 0; i < 48; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
}

// printVertical prints a vertical grid element
func printVertical() {
	fmt.Printf("|")
}

/* func getRandom() [][]int {
	// Store all available numbers from 2 to 2048
	arr := make([]int, 0)
	arr = append(arr, 0)
	for val := 2; val <= 2048; val *= 2 {
		arr = append(arr, val)
	}
	size := len(arr)

	// generate random numbers for init board
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	board := make([][]int, _rows)
	for i := 0; i < _rows; i++ {
		board[i] = make([]int, _cols)
	}
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			elem := arr[r.Int()%size]
			board[i][j] = elem
		}
	}
	return board
} */

func New() *board {
	matrix := make([][]int, _rows)
	for i := 0; i < _rows; i++ {
		matrix[i] = make([]int, _cols)
	}
	return &board{matrix: matrix}
}
