package medium_quickperm_permutations

/*
https://leetcode.com/problems/permutations/
*/
func permute(nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{nums}
	}

	permutation := make([]int, len(nums))
	copy(permutation, nums)
	permutations := [][]int{permutation}

	p := make([]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		p[i] = i
	}

	for i := 1; i < len(nums); {
		p[i]--
		j := i % 2 * p[i]
		nums[i], nums[j] = nums[j], nums[i]
		permutation = make([]int, len(nums))
		copy(permutation, nums)
		permutations = append(permutations, permutation)
		i = 1
		for p[i] == 0 {
			p[i] = i
			i++
		}
	}

	return permutations
}
