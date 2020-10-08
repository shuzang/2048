package game

func (b *board) moveLeft() {
	for i := 0; i < _rows; i++ {
		old := b.board[i]
		b.board[i] = moveRow(old)
	}
}

func (b *board) moveRight() {
	b.Reverse()
	b.moveLeft()
	b.Reverse()
}

func (b *board) moveUp() {
	b.leftRotate90()
	b.moveLeft()
	b.rightRotate90()
}

func (b *board) moveDown() {
	b.rightRotate90()
	b.moveLeft()
	b.leftRotate90()
}

func (b *board) rightRotate90() {
	matrix := make([][]int, _rows)
	for i := 0; i < _rows; i++ {
		matrix[i] = make([]int, _cols)
	}
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			matrix[j][_cols-1-i] = b.board[i][j]
		}
	}
	b.board = matrix
}

func (b *board) leftRotate90() {
	matrix := make([][]int, _rows)
	for i := 0; i < _rows; i++ {
		matrix[i] = make([]int, _cols)
	}
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			matrix[_cols-1-j][i] = b.board[i][j]
		}
	}
	b.board = matrix
}

func (b *board) Reverse() {
	for i := 0; i < _rows; i++ {
		for j, k := 0, _cols-1; j < k; j, k = j+1, k-1 {
			b.board[i][j], b.board[i][k] = b.board[i][k], b.board[i][j]
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
