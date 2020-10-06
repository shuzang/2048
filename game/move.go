package game

type Dir int

const (
	UP Dir = iota
	DOWN
	LEFT
	RIGHT
	QUIT
	ERROR_KEY
)

func (b *board) moveLeft() {
	for i := 0; i < _rows; i++ {
		old := b.matrix[i]
		b.matrix[i] = moveRow(old)
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

func moveRow(row []int) []int {
	index := 0
	for i := 0; i < len(row); i++ {
		if row[i] != 0 {
			row[index], row[i] = row[i], row[index]
			index++
		}
	}
	for i := 0; i < len(row)-1; i++ {
		if row[i] == row[i+1] {
			row[i] += row[i+1]
			row[i+1] = 0
			i++
		}
	}
	index = 0
	for i := 0; i < len(row); i++ {
		if row[i] != 0 {
			row[index], row[i] = row[i], row[index]
			index++
		}
	}
	return row
}
