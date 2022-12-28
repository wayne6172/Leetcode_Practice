package findLength

func findLength(nums1 []int, nums2 []int) int {
	nums1Length := len(nums1)
	nums2Length := len(nums2)

	dp := make([][]int, nums1Length+1)

	for i := range dp {
		dp[i] = make([]int, nums2Length+1)
	}

	ans := 0
	for i := nums1Length - 1; i >= 0; i-- {
		for j := nums2Length - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
			}
		}
	}

	return ans
}
