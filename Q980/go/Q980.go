package Q980

func findStartAndCountNotWall(grid [][]int) (int, int, int, int, int) {
	var startX, startY int
	row := len(grid)
	col := len(grid[0])
	notWall := 0
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			if grid[x][y] == 1 {
				startX, startY = x, y
			}
			if grid[x][y] != -1 {
				notWall++
			}
		}
	}

	return startX, startY, row, col, notWall
}

func uniquePathsIII(grid [][]int) int {
	startX, startY, row, col, notWalls := findStartAndCountNotWall(grid)

	var dfs func([][]int, int, int)

	length := 0
	ans := 0
	dfs = func(grid [][]int, x int, y int) {
		if x >= 0 && y >= 0 && x < row && y < col && grid[x][y] != -1 {
			length++

			if grid[x][y] == 2 && length == notWalls {
				ans++
			} else if grid[x][y] != 2 {
				grid[x][y] = -1
				dfs(grid, x+1, y)
				dfs(grid, x, y+1)
				dfs(grid, x-1, y)
				dfs(grid, x, y-1)

				grid[x][y] = 0
			}

			length--
		}
	}

	dfs(grid, startX, startY)

	return ans
}
