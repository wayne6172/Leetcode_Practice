package contest

func countPair(len int) int {
	if len < 0 {
		return 0
	}

	return (len * (len - 1)) / 2
}

func countGood(nums []int, k int) int64 {
	exist := map[int]int{}
	exist_pair := 0
	ans := int64(0)

	left := 0
	for _, value := range nums {
		existValue := exist[value]
		exist_pair += countPair(existValue+1) - countPair(existValue)
		exist[value]++

		for {
			x := exist[nums[left]]
			d := countPair(x) - countPair(x-1)
			if exist_pair-d < k {
				break
			}
			exist[nums[left]]--
			exist_pair -= d
			left++
		}

		if exist_pair >= k {
			ans += int64(left) + 1
		}
	}

	return ans
}

// [2,3,1,3,2,3,3,3,1,1,3,2,2,2]
// 3,1,6
//
