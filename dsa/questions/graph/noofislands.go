package graph

func numIslands(grid [][]byte) (noOfIslands int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	row := len(grid)
	column := len(grid[0])
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, column)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				noOfIslands++
				bfs(grid, visited, i, j)
			}
		}
	}
	return noOfIslands
}

func bfs(grid [][]byte, visited [][]bool, rowIndex int, columnIndex int) {
	type queuType struct {
		i int
		j int
	}
	visited[rowIndex][columnIndex] = true
	queue := []queuType{{rowIndex, columnIndex}}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		// Pop
		nodeToBeExplore := queue[0]
		queue = queue[1:]

		for _, d := range directions {
			// Validate index
			vR := nodeToBeExplore.i + d[0]
			vC := nodeToBeExplore.j + d[1]
			if vR >= 0 && vR < len(grid) && vC >= 0 && vC < len(grid[0]) {
				if !visited[vR][vC] && grid[vR][vC] == '1' {
					visited[vR][vC] = true
					queue = append(queue, queuType{vR, vC})
				}
			}
		}
	}
}
