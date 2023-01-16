package Q2531

func isItPossible(word1 string, word2 string) bool {
	mask1 := make([]int, 26)
	mask2 := make([]int, 26)
	word1_cat := 0
	word2_cat := 0

	for _, c := range word1 {
		if mask1[c-'a'] == 0 {
			word1_cat++
		}
		mask1[c-'a']++
	}

	for _, c := range word2 {
		if mask2[c-'a'] == 0 {
			word2_cat++
		}
		mask2[c-'a']++
	}

	for i := 0; i < 26; i++ {
		if mask1[i] > 0 {
			for j := 0; j < 26; j++ {
				if mask2[j] > 0 {
					mask1[i]--
					mask2[i]++

					mask1[j]++
					mask2[j]--

					c1, c2 := 0, 0
					for t := 0; t < 26; t++ {
						if mask1[t] > 0 {
							c1++
						}
						if mask2[t] > 0 {
							c2++
						}
					}

					if c1 == c2 {
						return true
					}

					mask1[i]++
					mask2[i]--

					mask1[j]--
					mask2[j]++
				}
			}
		}
	}

	return false
}
