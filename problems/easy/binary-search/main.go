package easy_binary_search

/**
https://leetcode.com/problems/binary-search/
*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	minPos := 0
	maxPos := len(nums) - 1
	curPos := maxPos / 2
	for nums[curPos] != target {
		if nums[curPos] > target {
			maxPos = curPos
			curPos = minPos + (maxPos-minPos)/2
		} else {
			minPos = curPos
			curPos = minPos + (maxPos-minPos+1)/2
		}

		if maxPos-minPos <= 1 {
			if nums[minPos] == target {
				return minPos
			}

			if nums[maxPos] == target {
				return maxPos
			}

			return -1
		}
	}

	return curPos
}
