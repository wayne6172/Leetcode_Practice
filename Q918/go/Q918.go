package Q918

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func maxSubarraySumCircular(nums []int) int {
	total, maxSum, curMax, minSum, curMin := 0, nums[0], 0, nums[0], 0
	for _, num := range nums {
		curMax = max(curMax+num, num)
		maxSum = max(maxSum, curMax)
		curMin = min(curMin+num, num)
		minSum = min(minSum, curMin)
		total += num
	}

	if maxSum > 0 {
		return max(maxSum, total-minSum)
	} else {
		return maxSum
	}
}

// 5 -3 5
