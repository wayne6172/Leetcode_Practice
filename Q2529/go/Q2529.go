package Q2529

func maximumCount(nums []int) int {
	x, y := 0, 0
	for _, value := range nums {
		if value > 0 {
			x++
		} else if value < 0 {
			y++
		}
	}

	if x > y {
		return x
	} else {
		return y
	}
}
