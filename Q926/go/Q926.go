package Q926

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func minFlipsMonoIncr(s string) int {

	is0 := make([]int, len(s)+1)
	is1 := make([]int, len(s)+1)

	for i, c := range s {
		index := i + 1
		if c == '0' {
			is0[index] = is0[index-1]
			is1[index] = min(is1[index-1]+1, is0[index-1]+1)
		} else {
			is1[index] = min(is0[index-1], is1[index-1])
			is0[index] = is0[index-1] + 1
		}
	}

	return min(is0[len(s)], is1[len(s)])
}
