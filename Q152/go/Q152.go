package Q152

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	neg := 0
	pos := 0
	ans := -1

	for _, v := range nums {
		if v > 0 {
			neg = neg * v
			if pos == 0 {
				pos = v
			} else {
				pos = pos * v
			}
		} else if v < 0 {
			prePos := pos
			pos = neg * v
			if prePos == 0 {
				neg = v
			} else {
				neg = prePos * v
			}
		} else {
			pos, neg = 0, 0
		}

		ans = max(ans, pos)
	}

	return ans
}

//   2 3 -2  4   -7  -13   0 -2 5
// 0 0 0 -12 -48 -28 -4368 0 -2 -10
// 0 2 6 0   4   336 364   0 0  5
