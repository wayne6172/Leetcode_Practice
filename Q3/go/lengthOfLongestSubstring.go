package lengthOfLongestSubstring

// 16 ms (43.16%), 2.8 MB (81.71%), 2022/3/5 12:53
func lengthOfLongestSubstring(s string) int {
	i, j, ans := 0, 0, 0
	alphaMap := make(map[byte]int)

	for ; j < len(s); j++ {
		if _, exist := alphaMap[s[j]]; exist {
			for ; i <= alphaMap[s[j]]; i++ {

				delete(alphaMap, s[i])
			}
		}

		ans = max(ans, j-i+1)
		alphaMap[s[j]] = j
	}

	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
