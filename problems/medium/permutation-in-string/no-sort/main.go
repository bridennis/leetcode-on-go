package medium_permutation_in_string_no_sort

/**
https://leetcode.com/problems/permutation-in-string/
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) || len(s1) == 0 || len(s2) == 0 {
		return false
	}

	s1Q := make(map[byte]int, 26) // char frequency map
	for i := range s1 {
		s1Q[s1[i]]++
	}

	s2Q := make(map[byte]int, 26) // char frequency map
	for j := range s1 {
		s2Q[s2[j]]++
	}

	for i := len(s1); i < len(s2); i++ { // sliding window
		if intersects(s1Q, s2Q) {
			return true
		}
		s2Q[s2[i]]++
		s2Q[s2[i-len(s1)]]--
	}

	return intersects(s1Q, s2Q)
}

func intersects(q1 map[byte]int, q2 map[byte]int) bool {
	for k, v := range q1 {
		if v2, ok := q2[k]; !ok || v2 != v {
			return false
		}
	}

	return true
}
