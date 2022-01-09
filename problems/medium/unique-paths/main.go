package medium_unique_paths

/*
https://leetcode.com/problems/unique-paths/
*/
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	if m < n {
		n, m = m, n
	}

	row := make([]int, n)
	for i := range row {
		row[i] = 1
	}

	for i := 2; i <= m; i++ {
		for j := 1; j < n; j++ {
			row[j] = row[j] + row[j-1]
		}
	}

	return row[n-1]
}
