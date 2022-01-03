package medium_permutations

/*
https://leetcode.com/problems/permutations/
Non-Recursive Heap’s Algorithm
*/
func permute(nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{nums}
	}

	permutation := make([]int, len(nums))
	copy(permutation, nums)
	permutations := [][]int{permutation}

	i := 0
	pos := make([]int, len(nums))
	for i < len(nums) {
		if pos[i] < i {
			if i%2 == 0 {
				nums[0], nums[i] = nums[i], nums[0]
			} else {
				nums[pos[i]], nums[i] = nums[i], nums[pos[i]]
			}
			permutation = make([]int, len(nums))
			copy(permutation, nums)
			permutations = append(permutations, permutation)
			pos[i] += 1
			i = 0
		} else {
			pos[i] = 0
			i++
		}
	}

	return permutations
}
