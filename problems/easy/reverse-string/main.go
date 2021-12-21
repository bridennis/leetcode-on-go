package easy_reverse_string

/*
https://leetcode.com/problems/reverse-string/
*/
func reverseString(s []byte) {
	if len(s) == 1 {
		return
	}

	for i := 0; i < len(s); i++ {
		j := len(s) - 1 - i
		if j <= i {
			break
		}
		s[i], s[j] = s[j], s[i]
	}
}
