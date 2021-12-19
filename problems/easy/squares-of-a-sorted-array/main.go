package easy_squares_of_a_sorted_array

/**
https://leetcode.com/problems/squares-of-a-sorted-array/
*/
func sortedSquares(nums []int) []int {
	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	}

	var sqn int
	sq := make([]int, 0, len(nums))
	for _, num := range nums {
		sqn = num * num
		for i, v := range sq {
			if v >= sqn {
				sq = append(sq[:i+1], sq[i:]...)
				sq[i] = sqn
				goto Next
			}
		}
		sq = append(sq, sqn)
	Next:
	}

	return sq
}
