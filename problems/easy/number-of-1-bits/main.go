package easy_number_of_1_bits

/*
https://leetcode.com/problems/number-of-1-bits/
*/
func hammingWeight(num uint32) int {
	weight := 0
	for num > 0 {
		if num&1 == 1 {
			weight++
		}
		num = num >> 1
	}
	return weight
}
