package medium_combinations

/*
https://leetcode.com/problems/combinations/
*/
func combine(n int, k int) [][]int {
	var combines [][]int

	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}

	max := pow(2, n)
	bits := make([]int, n)
	for i := 1; i < max; i++ {
		if countBits(i) == k {
			numToBitsArr(i, bits)
			combines = append(combines, combineFromBits(arr, bits, k))
		}
	}

	return combines
}

func numToBitsArr(num int, bits []int) []int {
	for i := 0; num > 0; i++ {
		if num%2 == 0 {
			bits[len(bits)-1-i] = 0
		} else {
			bits[len(bits)-1-i] = 1
		}
		num /= 2
	}

	return bits
}

func countBits(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}

func combineFromBits(arr []int, bits []int, k int) []int {
	combine := make([]int, 0, k)
	for i := range bits {
		if bits[len(bits)-1-i] == 1 {
			combine = append(combine, arr[len(arr)-1-i])
		}
	}

	return combine
}

func pow(n int, p int) int {
	pow := 1
	for ; p > 0; p-- {
		pow *= n
	}

	return pow
}
