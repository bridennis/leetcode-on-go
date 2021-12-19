package easy_search_insert_position

/**
https://leetcode.com/problems/search-insert-position/
*/
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	if target < nums[0] {
		return 0
	}

	minPos := 0
	maxPos := len(nums) - 1
	curPos := maxPos / 2
	for {
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

			return maxPos
		}
	}
}
