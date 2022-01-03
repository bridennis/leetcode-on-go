package medium_house_robber

/*
https://leetcode.com/problems/house-robber/
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var accum, val2, maxAmount int
	for i := 0; i < len(nums); i++ {
		if accum+nums[i] > val2 {
			maxAmount = accum + nums[i]
		} else {
			maxAmount = val2
		}

		accum, val2 = val2, maxAmount
	}

	return maxAmount
}
