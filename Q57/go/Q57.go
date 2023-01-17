package Q57

func isInRange(range1 []int, range2 []int) bool {
	if range1[0] < range2[0] {
		if range1[1] < range2[0] {
			return false
		}
	} else {
		if range1[0] > range2[1] {
			return false
		}
	}

	return true
}

//
//             6 --------- 9
// 2-------5

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a int, b int) int {
	if a < b {
		return b
	}

	return a
}

func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	first, end := -1, -1

	index := 0
	for _, interval := range intervals {
		if end == -1 && isInRange(interval, newInterval) {
			if first == -1 {
				first = index
			}
		} else {
			if first != -1 {
				break
			}

			if newInterval[1] < interval[0] {
				break
			}

			ans = append(ans, interval)
		}

		index++
	}

	if first != -1 {
		end = index - 1
		x := min(intervals[first][0], newInterval[0])
		y := max(intervals[end][1], newInterval[1])
		ans = append(ans, []int{x, y})

		ans = append(ans, intervals[index:]...)
	} else {
		ans = append(ans, newInterval)
		ans = append(ans, intervals[index:]...)
	}

	return ans
}
