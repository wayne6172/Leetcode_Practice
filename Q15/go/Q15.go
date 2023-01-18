package Q15

import "sort"

func threeSum(nums []int) [][]int {
	var ans [][]int
	length := len(nums)

	sort.Ints(nums)
	preI, preJ := 10000000, 10000000
	for i := 0; i < length; i++ {
		if preI == nums[i] {
			continue
		}

		for j := i + 1; j < length; j++ {
			if preJ == nums[j] {
				continue
			}

			t := nums[(j + 1):]
			k := sort.SearchInts(t, -(nums[i] + nums[j]))

			if k < len(t) && t[k] == -(nums[i]+nums[j]) {
				ans = append(ans, []int{nums[i], nums[j], t[k]})
			}

			preJ = nums[j]
		}

		preI = nums[i]
	}

	return ans
}
