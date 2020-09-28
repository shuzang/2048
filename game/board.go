package game

import (
	"fmt"
	"math/rand"
	"time"
)

// game board size
const rows, cols = 4, 4

type Board interface {
	Display()
}

type board struct {
	matrix [][]int
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
func (b board) Display() {
	b.matrix = getRandom()
	printHorizontal()
	for i := 0; i < len(b.matrix); i++ {
		for j := 0; j < len(b.matrix[0]); j++ {
			if b.matrix[i][j] == 0 {
				fmt.Printf("%6s", "")
			} else {
				fmt.Printf("%6d", b.matrix[i][j])
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

func printHorizontal() {
	for i := 0; i < 48; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
}

func printVertical() {
	fmt.Printf("|")
}

func getRandom() [][]int {
	// Store all available numbers from 2 to 2048
	arr := make([]int, 0)
	arr = append(arr, 0)
	for val := 2; val <= 2048; val *= 2 {
		arr = append(arr, val)
	}
	size := len(arr)

	// generate random numbers for init board
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	board := make([][]int, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			elem := arr[r.Int()%size]
			board[i][j] = elem
		}
	}
	return board
}

func New() Board {
	return &board{}
}
