package Q49

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	m := map[string]([]string){}
	for _, str := range strs {
		t := SortString(str)
		m[t] = append(m[t], str)
	}

	ans := [][]string{}
	for _, value := range m {
		ans = append(ans, value)
	}

	return ans
}
