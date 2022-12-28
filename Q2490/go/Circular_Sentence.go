package Circular_Sentence

import "strings"

func isCircularSentence(sentence string) bool {

	sentences := strings.Split(sentence, " ")

	check := true
	length := len(sentences)
	for i := 0; i < length; i++ {
		s1 := sentences[i]
		s2 := sentences[(i+1)%length]

		s1_len := len(s1)

		if s1[s1_len-1] != s2[0] {
			check = false
			break
		}
	}

	return check
}
