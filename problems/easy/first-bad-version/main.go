package easy_first_bad_version

/**
https://leetcode.com/problems/first-bad-version/
*/
func firstBadVersion(n int) int {
	min := 1
	max := n
	var version int
	for {
		version = min + (max-min)/2
		if isBadVersion(version) {
			max = version
		} else {
			min = version
		}

		if max-min <= 1 {
			if isBadVersion(min) {
				return min
			} else {
				return max
			}
		}
	}
}

func isBadVersion(version int) bool {
	return version >= 4
}
