package Q93

import (
	"strconv"
	"strings"
)

func checkValid(s string) bool {
	if len(s) == 1 || s[0] != '0' {
		if res, _ := strconv.Atoi(s); res <= 255 {
			return true
		}
	}

	return false
}

func restoreIpAddresses(s string) []string {
	ans := []string{}

	var dfs func(string, int, []string)
	dfs = func(s string, now int, queue []string) {
		if now == 3 {
			if checkValid(s) {
				a := strings.Join(queue, ".")
				a += "." + s
				ans = append(ans, a)
			}
		} else {
			for i := 1; i <= len(s)-1 && i <= 3; i++ {
				f := s[0:i]
				if checkValid(f) {
					queue = append(queue, f)
					dfs(s[i:], now+1, queue)
					queue = queue[0 : len(queue)-1]
				}
			}
		}
	}
	dfs(s, 0, []string{})
	return ans
}
