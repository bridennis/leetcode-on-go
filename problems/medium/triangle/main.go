package medium_triangle

/*
https://leetcode.com/problems/triangle/
*/
func minimumTotal(triangle [][]int) int {
	minTotal := triangle[len(triangle)-1]
	min := 0
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if minTotal[j+1] < minTotal[j] {
				min = minTotal[j+1]
			} else {
				min = minTotal[j]
			}
			minTotal[j] = triangle[i][j] + min
		}
	}

	return minTotal[0]
}
