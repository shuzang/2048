package game

type Dir int

const (
	UP Dir = iota
	DOWN
	LEFT
	RIGHT
	NO_DIR
)

func (b *board) moveLeft() {
	for i := 0; i < _rows; i++ {
		old := b.matrix[i]
		b.matrix[i] = moveRow(old)
		//fmt.Printf("updated row is: %v || old row is: %v\n", b.matrix[i], old)
	}
}

func (b *board) moveRight() {
	b.reverse()
	b.moveLeft()
	b.reverse()
}

func (b *board) moveUp() {
	b.rightRotate90()
	b.moveLeft()
	b.leftRotate90()
}

func (b *board) moveDown() {
	b.leftRotate90()
	b.moveLeft()
	b.rightRotate90()
}

func (b *board) rightRotate90() {
	res := make([][]int, _rows)
	for i := 0; i < _rows; i++ {
		res[i] = make([]int, _cols)
	}
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			res[_cols-j-1][i] = b.matrix[i][j]
		}
	}
	b.matrix = res
}

func (b *board) leftRotate90() {
	res := make([][]int, _rows)
	for i := 0; i < _rows; i++ {
		res[i] = make([]int, _cols)
	}
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			res[j][_rows-i-1] = b.matrix[i][j]
		}
	}
	b.matrix = res
}

func (b *board) reverse() {
	for i := 0; i < _rows; i++ {
		for j, k := 0, _cols-1; j < k; j, k = j+1, k-1 {
			b.matrix[i][j], b.matrix[i][k] = b.matrix[i][k], b.matrix[i][j]
		}
	}
}

func moveRow(elems []int) []int {
	index := 0
	for i := 0; i < _cols; i++ {
		if elems[i] != 0 {
			elems[index], elems[i] = elems[i], elems[index]
			index++
		}
	}
	return mergeElements(elems)
}

func mergeElements(arr []int) []int {
	newArr := make([]int, len(arr))
	newArr[0] = arr[0]
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == newArr[index] {
			newArr[index] += arr[i]
		} else {
			index++
			newArr[index] = arr[i]
		}
	}
	return newArr
}
