package game

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

const (
	_rows, _cols = 4, 4
)

type Key int

const (
	UP Key = iota
	DOWN
	LEFT
	RIGHT
	QUIT
	ERROR_KEY
)

type board struct {
	board  [][]int
	nx, ny int
}

type Board interface {
	Display()
	AddElement()
	TakeInput() bool
	IsOver() bool
	CountScore() (int, int)
}

func (b *board) CountScore() (int, int) {
	total, max := 0, 0
	matrix := b.board
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
	blank := 0
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.board[i][j] == 0 {
				blank++
			}
		}
	}
	return blank == 0
}

func (b *board) TakeInput() bool {
	key, err := GetKeyStrokes()
	if err != nil {
		fmt.Printf(err.Error())
	}
	if key == ERROR_KEY {
		b.TakeInput()
	}
	switch key {
	case UP:
		b.moveUp()
	case DOWN:
		b.moveDown()
	case LEFT:
		b.moveLeft()
	case RIGHT:
		b.moveRight()
	case QUIT:
		fmt.Println("You press ESC, game exit!")
		return false
	}
	return true
}

func GetKeyStrokes() (Key, error) {
	char, key, err := keyboard.GetSingleKey()
	if err != nil {
		return ERROR_KEY, err
	}
	//fmt.Printf("You pressed: %c, key %X\r\n", char, key)
	if int(char) == 0 {
		switch key {
		case keyboard.KeyArrowUp:
			return UP, nil
		case keyboard.KeyArrowDown:
			return DOWN, nil
		case keyboard.KeyArrowLeft:
			return LEFT, nil
		case keyboard.KeyArrowRight:
			return RIGHT, nil
		case keyboard.KeyEsc:
			return QUIT, nil
		default:
			return ERROR_KEY, errors.New("Key input error, please press again!")
		}
	} else {
		switch char {
		case 119:
			return UP, nil
		case 97:
			return LEFT, nil
		case 115:
			return DOWN, nil
		case 100:
			return RIGHT, nil
		default:
			return ERROR_KEY, errors.New("Key input error, please press again!")
		}
	}
}

func (b *board) AddElement() {
	rand.Seed(time.Now().UnixNano())
	// get random matrix index
	index := make([][2]int, 0)
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.board[i][j] == 0 {
				index = append(index, [2]int{i, j})
			}
		}
	}
	next := rand.Int() % len(index)
	nx, ny := index[next][0], index[next][1]
	// get random number
	var number int
	if rand.Int()%100 < 80 {
		number = 2
	} else {
		number = 4
	}
	// filling
	b.nx, b.ny = nx, ny
	b.board[nx][ny] = number
}

func (b *board) Display() {
	// clear screen, but only works on windows
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	c := color.New(color.FgCyan, color.Bold)
	printHorizontalLine()
	for i := 0; i < _rows; i++ {
		printVerticalLine()
		for j := 0; j < _cols; j++ {
			if b.board[i][j] == 0 {
				fmt.Printf("%7s", "")
			} else if i == b.nx && j == b.ny {
				c.Printf("%4d%3s", b.board[i][j], "")
			} else {
				fmt.Printf("%4d%3s", b.board[i][j], "")
			}

			printVerticalLine()
		}
		fmt.Println()
		printHorizontalLine()
	}

}

func printHorizontalLine() {
	for i := 0; i < 33; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
}

func printVerticalLine() {
	fmt.Printf("|")
}

/* func generate(matrix [][]int) [][]int {
	// Generate a number pool
	nums := make([]int, 0)
	nums = append(nums, 0)
	for i := 2; i <= 2048; i *= 2 {
		nums = append(nums, i)
	}

	// fill the matrix using random number
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			matrix[i][j] = nums[rand.Int()%len(nums)]
		}
	}

	return matrix
} */

func NewBoard() *board {
	matrix := make([][]int, _rows)
	for i := 0; i < _rows; i++ {
		matrix[i] = make([]int, _cols)
	}
	return &board{board: matrix}
}
