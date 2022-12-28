package Q55

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func canJump(nums []int) bool {
	l := len(nums)
	mask := make([]bool, l)

	max_arrive := 0
	mask[0] = true
	for index := range nums {
		if mask[index] && index+nums[index] > max_arrive {
			start := max_arrive - index
			for i := start + 1; i <= nums[index] && index+i < l; i++ {
				mask[index+i] = true
				max_arrive = max(max_arrive, index+i)
			}
		}
	}

	return mask[l-1]
}
