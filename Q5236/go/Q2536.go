package contest

func rangeAddQueries(n int, queries [][]int) [][]int {
	g := make([][]int, n)
	for i := 0; i < len(g); i++ {
		g[i] = make([]int, n)
	}

	for _, query := range queries {
		for i := query[0]; i <= query[2]; i++ {
			for j := query[1]; j <= query[3]; j++ {
				g[i][j]++
			}
		}
	}

	return g
}
