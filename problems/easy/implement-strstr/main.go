package easy_implement_strstr

/*
https://leetcode.com/problems/implement-strstr/
*/
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] && haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
