package countBits

// 9ms (35.58%) 4.5 MB (68.54%), O(n), 2022/3/5 13:35
// 此題可O(n logn)， 但可發想為 O(n)
func countBits(n int) []int {
	switch n {
	case 0:
		return []int{0}
	case 1:
		return []int{0, 1}
	}

	var (
		ans  []int = make([]int, n+1)
		from int   = 0
		now  int   = 2
		x    int   = 2
	)

	ans[0] = 0
	ans[1] = 1

	for now <= n {
		from = 0
		x = x * 2
		for ; now <= n && now < x; now++ {
			ans[now] = 1 + ans[from]
			from++
		}
	}

	return ans
}
