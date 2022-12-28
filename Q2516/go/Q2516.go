package Q2516

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func takeCharacters(s string, k int) int {
	l := len(s)
	mask := []int{0, 0, 0}
	for _, c := range s {
		mask[c-'a']++
	}

	ans := 0x0fffffff
	g := l - 1
	for i := l; i >= 0; i-- {

		for g >= i && !(mask[0] >= k && mask[1] >= k && mask[2] >= k) {
			mask[s[g]-'a']++
			g--
		}

		if mask[0] >= k && mask[1] >= k && mask[2] >= k {
			ans = min(ans, i+l-g)
		}

		if i > 0 {
			mask[s[i-1]-'a']--
		}
	}

	if ans == 0x0fffffff {
		return -1
	}

	return ans - 1
}
