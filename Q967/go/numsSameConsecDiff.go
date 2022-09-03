package numsSameConsecDiff

func numsSameConsecDiff(n int, k int) []int {

	ans := make([]int, 0)

	var dfs func(nowDigitLevel int, n int, digitNumber int, k int, nowValue int)
	dfs = func(nowDigitLevel int, n int, digitNumber int, k int, nowValue int) {
		if 0 <= digitNumber && digitNumber < 10 {
			if nowDigitLevel < n-1 {
				if k == 0 {
					dfs(nowDigitLevel+1, n, digitNumber, k, nowValue*10+digitNumber)
				} else {
					dfs(nowDigitLevel+1, n, digitNumber+k, k, nowValue*10+digitNumber)
					dfs(nowDigitLevel+1, n, digitNumber-k, k, nowValue*10+digitNumber)
				}
			} else {
				ans = append(ans, nowValue*10+digitNumber)
			}
		}
	}

	for i := 1; i < 10; i++ {
		dfs(0, n, i, k, 0)
	}

	return ans
}
