package medium_longest_substring_without_repeating_characters

/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	substring := ""
	maxLen := 1
	for _, c := range s {
		substring = nextSubstring(substring, c)
		if len(substring) > maxLen {
			maxLen = len(substring)
		}
	}

	return maxLen
}

func nextSubstring(s string, char int32) string {
	for i, c := range s {
		if c == char {
			s = s[i+1:]
			break
		}
	}
	return s + string(char)
}
