package Q2491

import (
	"sort"
)

func dividePlayers(skill []int) int64 {
	sort.Slice(skill, func(i, j int) bool {
		return skill[i] > skill[j]
	})

	length := len(skill)
	sum := skill[0] + skill[length-1]

	var ans int64 = 0
	for i, j := 0, length-1; i < j; {
		if skill[i]+skill[j] != sum {
			return -1
		}

		ans += int64(skill[i] * skill[j])

		i++
		j--
	}

	return ans
}
