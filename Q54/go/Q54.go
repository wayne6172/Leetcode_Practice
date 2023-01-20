package Q54

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])

	direct := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	ans := []int{}
	x, i, j, dir := 0, 0, -1, 0
	for x < rows*cols {
		var newI, newJ int = i, j
		for {
			newI = i + direct[dir][0]
			newJ = j + direct[dir][1]

			if 0 <= newI && newI < rows && 0 <= newJ && newJ < cols && matrix[newI][newJ] != -101 {
				break
			}
			dir = (dir + 1) % 4
		}

		ans = append(ans, matrix[newI][newJ])
		matrix[newI][newJ] = -101
		i = newI
		j = newJ
		x++
	}

	return ans
}
