package backtracking

func exist(board [][]byte, word string) bool {
	row := len(board)
	column := len(board[0])

	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, column)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if board[i][j] == word[0] {
				if wordSearch(board, i, j, &word, 0, visited) {
					return true
				}
			}
		}
	}
	return false
}

func wordSearch(board [][]byte, rowNo int, colNo int, word *string, charNoToBeSearch int, visited [][]bool) bool {
	row := len(board)
	col := len(board[0])

	// Base condition
	if charNoToBeSearch == len(*word) {
		return true
	}

	// Base condition
	if rowNo < 0 || rowNo >= row || colNo < 0 || colNo >= col || charNoToBeSearch >= len(*word) || visited[rowNo][colNo] || board[rowNo][colNo] != (*word)[charNoToBeSearch] {
		return false
	}
	visited[rowNo][colNo] = true

	// 0,1 - 0,-1 - 1,0, -1,0
	// Find in the four directions
	dir1 := wordSearch(board, rowNo, colNo+1, word, charNoToBeSearch+1, visited)
	dir2 := wordSearch(board, rowNo, colNo-1, word, charNoToBeSearch+1, visited)
	dir3 := wordSearch(board, rowNo+1, colNo, word, charNoToBeSearch+1, visited)
	dir4 := wordSearch(board, rowNo-1, colNo, word, charNoToBeSearch+1, visited)

	// Backtrack: Unmark the cell as visited for (other paths)
	visited[rowNo][colNo] = false

	return dir1 || dir2 || dir3 || dir4
}
