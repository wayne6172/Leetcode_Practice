package Q56

import "sort"

func isIn(interval []int, l int, r int) bool {
	if interval[0] < l {
		return l <= interval[1]
	} else {
		return interval[0] <= r
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{}
	l, r := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if isIn(intervals[i], l, r) {
			l = min(l, intervals[i][0])
			r = max(r, intervals[i][1])
		} else {
			ans = append(ans, []int{l, r})
			l = intervals[i][0]
			r = intervals[i][1]
		}
	}
	ans = append(ans, []int{l, r})
	return ans
}
