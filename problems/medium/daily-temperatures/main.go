package easy_daily_temperatures

/*
https://leetcode.com/problems/daily-temperatures/
*/
func dailyTemperatures(temperatures []int) []int {
	warmerAfterDays := make([]int, len(temperatures))
	stack := []int{len(temperatures) - 1} // stack of maximal temp. indexes
	for i := len(temperatures) - 2; i > -1; i-- {
		j := len(stack) - 1
		for ; j > -1; j-- {
			if temperatures[stack[j]] > temperatures[i] {
				stack = append(stack[:j+1], i)
				warmerAfterDays[i] = stack[j] - i
				goto Next
			}
		}
		stack = []int{i}
		warmerAfterDays[i] = 0
	Next:
	}

	return warmerAfterDays
}
