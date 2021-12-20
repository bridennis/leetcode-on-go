package easy_two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	firstNumbers := make(map[int]int)

	for i, n := range numbers {
		if j, ok := firstNumbers[target-n]; ok {
			return []int{j + 1, i + 1}
		}

		firstNumbers[n] = i
	}

	return nil
}
