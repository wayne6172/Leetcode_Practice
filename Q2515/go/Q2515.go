package Q251t

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x int, y int) int {
	if x > y {
		return y
	}

	return x
}

func closetTarget(words []string, target string, startIndex int) int {

	q := []int{}

	for index, word := range words {
		if word == target {
			q = append(q, index)
		}
	}

	short := 1000

	if len(q) == 0 {
		return -1
	}

	for _, val := range q {

		path1 := abs(startIndex - val)
		path2 := len(words) - path1

		short = min(short, path1)
		short = min(short, path2)
	}

	return short
}
