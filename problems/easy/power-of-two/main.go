package easy_power_of_two

/*
https://leetcode.com/problems/power-of-two/
*/
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
