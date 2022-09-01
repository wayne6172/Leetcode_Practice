package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	var ans [][]int
	length := len(nums)

	sort.Ints(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			t := nums[(j + 1):]
			k := sort.SearchInts(t, -(nums[i] + nums[j]))

			if k != len(t) {
				ans = append(ans, []int{nums[i], nums[j], t[k]})
			}
		}
	}

	return ans
}
