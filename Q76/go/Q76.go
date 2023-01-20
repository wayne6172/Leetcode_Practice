package Q76

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func minWindow(s string, t string) string {
	m := map[rune]int{}
	category := 0
	for _, c := range t {
		if _, isExist := m[c]; !isExist {
			category++
		}

		m[c]++
	}

	minL := 1000000
	ansR, ansL := -1, -1
	l := 0
	for r, c := range s {
		if n, isExist := m[c]; isExist {
			m[c]--

			if n-1 == 0 {
				category--
			}
		}

		for category == 0 {
			if r-l+1 < minL {
				ansL = l
				ansR = r
				minL = r - l + 1
			}

			if n, isExist := m[rune(s[l])]; isExist {
				m[rune(s[l])]++
				if n+1 > 0 {
					category++
				}
			}
			l++
		}
	}

	if ansL == -1 {
		return ""
	}

	return s[ansL : ansR+1]
}
