package easy_reverse_words_in_a_string_iii

/*
https://leetcode.com/problems/reverse-words-in-a-string-iii/
*/
func reverseWords(s string) string {
	if len(s) == 1 {
		return s
	}

	sArr := []byte(s)
	space := []byte(" ")[0]

	j := 0
	i := 0
	for i = 0; i < len(sArr); i++ {
		if sArr[i] == space {
			reverseWord(sArr, i, j)
			j = i + 1
		}
	}

	if j < i {
		reverseWord(sArr, i, j)
	}

	return string(sArr)
}

func reverseWord(s []byte, i int, j int) {
	wordLen := i - j
	for k := 0; k < wordLen; k++ {
		ii := j + k
		jj := j + wordLen - 1 - k
		if jj <= ii {
			break
		}
		s[ii], s[jj] = s[jj], s[ii]
	}
}
