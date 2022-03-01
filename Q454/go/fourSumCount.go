package fourSumCount

func neg(x int) int {
	return x * -1
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sums1 := make(map[int]int)
	sums2 := make(map[int]int)

	for index1, _ := range nums1 {
		for _, value2 := range nums2 {
			sums1[nums1[index1]+value2]++
		}
	}

	for index3, _ := range nums3 {
		for _, value4 := range nums4 {
			sums2[nums3[index3]+value4]++
		}
	}

	ans := 0
	for indexS1 := range sums1 {
		x := sums1[indexS1]
		y := sums2[neg(indexS1)]

		ans += x * y
	}

	return ans
}
