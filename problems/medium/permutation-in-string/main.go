package medium_permutation_in_string

import (
	"sort"
)

/**
https://leetcode.com/problems/permutation-in-string/
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}

	s := sortCharsInString([]byte(s1))
	for i := 0; i < len(s2)-len(s1)+1; i++ {
		if s == sortCharsInString([]byte(s2[i:i+len(s1)])) {
			return true
		}
	}

	return false
}

func sortCharsInString(chars []byte) string {
	sort.Slice(chars, func(i int, j int) bool { return chars[i] < chars[j] })
	return string(chars)
}
