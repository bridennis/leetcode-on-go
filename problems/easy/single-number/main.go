package easy_single_number

/*
https://leetcode.com/problems/single-number/
*/
func singleNumber(nums []int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		n ^= nums[i]
	}

	return n
}
