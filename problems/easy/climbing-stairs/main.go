package easy_climbing_stairs

/*
https://leetcode.com/problems/climbing-stairs/
*/
func climbStairs(n int) int {
	pre, curr := 1, 1
	for i := 1; i < n; i++ {
		curr, pre = curr+pre, curr
	}

	return curr
}
