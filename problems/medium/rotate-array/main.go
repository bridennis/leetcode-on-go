package medium_rotate_array

/**
https://leetcode.com/problems/rotate-array/
*/
func rotate(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}

	for i, r := range append(nums[len(nums)-k:], nums[:len(nums)-k]...) {
		nums[i] = r
	}
}
