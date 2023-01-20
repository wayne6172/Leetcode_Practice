package Q39

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] > candidates[j]
	})

	var dfs func(int, int, []int)

	ans := [][]int{}
	q := []int{}
	dfs = func(now int, t int, queue []int) {
		if t == 0 {
			ans = append(ans, queue)
		} else {
			for i := now; i < len(candidates); i++ {
				if t >= candidates[i] {
					queue = append(queue, candidates[i])
					dfs(i, t-candidates[i], queue)
					queue = queue[0 : len(queue)-1]
				}
			}
		}
	}

	dfs(0, target, q)
	return ans
}
