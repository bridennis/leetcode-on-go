package easy_move_zeroes

/**
https://leetcode.com/problems/move-zeroes/
*/
func moveZeroes(nums []int) {
	nonZeroPos := 0 // First non-zero position
	for cur := 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[nonZeroPos], nums[cur] = nums[cur], nums[nonZeroPos]
			nonZeroPos++
		}
	}
}
