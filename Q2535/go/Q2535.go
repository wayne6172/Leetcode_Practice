package Q2535

func getDigitSum(value int) int {
	a := 0
	for value != 0 {
		a += value % 10
		value /= 10
	}

	return a
}

func differenceOfSum(nums []int) int {

	x, y := 0, 0
	for _, value := range nums {
		x += getDigitSum(value)
		y += value
	}

	if x > y {
		return x - y
	}

	return y - x
}
