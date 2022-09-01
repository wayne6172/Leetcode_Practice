package numIslands

func fixMap(grid [][]byte, i int, j int, row int, col int) {

	if i < 0 || j < 0 || i >= row || j >= col {
		return
	}

	if grid[i][j] == '1' {
		grid[i][j] = '0'

		fixMap(grid, i+1, j, row, col)
		fixMap(grid, i, j+1, row, col)
		fixMap(grid, i-1, j, row, col)
		fixMap(grid, i, j-1, row, col)
	}
}

func numIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])

	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				fixMap(grid, i, j, row, col)
				count += 1
			}
		}
	}

	return count
}
