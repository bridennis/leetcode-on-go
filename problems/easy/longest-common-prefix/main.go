package easy_longest_common_prefix

/**
https://leetcode.com/problems/longest-common-prefix/
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	longestPrefix := ""
HorizontalScan:
	for offset := 0; ; offset++ {
		for i := range strs {
			if offset >= len(strs[i]) || (i > 0 && strs[i][offset] != strs[i-1][offset]) {
				break HorizontalScan
			}

			if i == len(strs)-1 {
				longestPrefix += string(strs[i][offset])
			}
		}
	}

	return longestPrefix
}
