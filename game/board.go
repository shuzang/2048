package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// game board size
const rows, cols = 4, 4
const _clearScreenSequence = "\033[H\033[2J"

type Board interface {
	Display()
	AddElement()
	TakeInput()
}

type board struct {
	matrix [][]int
}

func (b *board) TakeInput() {
	/* 	var char rune
	   	fmt.Scanf("%c", &char)
	   	fmt.Printf("keyboar input is: %v\n", char) */
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	switch input[0] {
	case 'a', 37:
		b.moveLeft()
	case 'd', 39:
		b.moveRight()
	case 'w', 38:
		b.moveUp()
	case 's', 40:
		b.moveDown()
	}
	fmt.Printf("Input char is: %v\n", input[0])
}

func (b *board) moveLeft() {
	for i := 0; i < rows; i++ {
		old := b.matrix[i]
		b.matrix[i] = moveRow(old)
		fmt.Printf("updated row is: %v || old row is: %v\n", b.matrix[i], old)
	}
}

func (b *board) moveRight() {
	b.reverse()
	b.moveLeft()
	b.reverse()
}

func (b *board) moveDown() {
	b.rightRotate90()
	b.moveLeft()
	b.leftRotate90()
}

func (b *board) moveUp() {
	b.leftRotate90()
	b.moveLeft()
	b.rightRotate90()
}

func (b *board) rightRotate90() {
	res := make([][]int, rows)
	for i := 0; i < rows; i++ {
		res[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res[cols-j-1][i] = b.matrix[i][j]
		}
	}
	b.matrix = res
}

func (b *board) leftRotate90() {
	res := make([][]int, rows)
	for i := 0; i < rows; i++ {
		res[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res[j][rows-i-1] = b.matrix[i][j]
		}
	}
	b.matrix = res
}

func (b *board) reverse() {
	for i := 0; i < rows; i++ {
		for j, k := 0, cols; j < k; j, k = j+1, k-1 {
			b.matrix[i][j], b.matrix[i][k] = b.matrix[i][k], b.matrix[i][j]
		}
	}
}

func moveRow(elems []int) []int {
	index := 0
	for i := 0; i < cols; i++ {
		if elems[i] != 0 {
			elems[index], elems[i] = elems[i], elems[index]
			index++
		}
	}
	return mergeElements(elems)
}

func mergeElements(arr []int) []int {
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[index] {
			arr[index] += arr[i]
		} else {
			index++
			arr[index] = arr[i]
		}
		arr[i] = 0
	}
	return arr
}

func (b *board) AddElement() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// generate new element
	val := r.Int() % 100
	if val < 70 {
		val = 2
	} else {
		val = 4
	}

	// count empty postions
	emptyCount := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if b.matrix[i][j] == 0 {
				emptyCount++
			}
		}
	}

	// generate the element position to be filled
	elementCount := r.Int()%emptyCount + 1
	index := 0

	// fill
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if b.matrix[i][j] == 0 {
				index++
				if index == elementCount {
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
	//b.matrix = getRandom()
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

func New() *board {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}
	return &board{matrix}
}
