package easy_letter_case_permutation

/*
https://leetcode.com/problems/letter-case-permutation/
 */
func letterCasePermutation(s string) []string {
	sByte := []byte(s)
	sCase := make([]int, 0) // letter case: 0 - lower, 1 - upper
	for i, value := range sByte {
		if upperCaseLetter(value) { // A-Z
			sCase = append(sCase, 0)
			sByte[i] += 32
		} else if lowerCaseLetter(value) { // a-z
			sCase = append(sCase, 0)
		}
	}

	permutations := make([]string, 0, pow(2, len(sCase)))
	permutations = append(permutations, string(sByte))

	bottom := len(sCase) - 1
	for i := 0; i < len(sCase); i++ {
		sCase[bottom-i] = 1
		for j := 0; j < i; j++ {
			sCase[bottom-j] = 0
		}
		permutations = append(permutations, casePermutation(sCase, sByte))

		for j := 1; j < pow(2, i); j++ {
			carry := j
			for k := 0; carry > 0; k++ {
				if carry%2 == 0 {
					sCase[bottom-k] = 0
				} else {
					sCase[bottom-k] = 1
				}
				carry /= 2
			}
			permutations = append(permutations, casePermutation(sCase, sByte))
		}
	}

	return permutations
}

func lowerCaseLetter(value byte) bool {
	return value >= 97 && value <= 122
}

func upperCaseLetter(value byte) bool {
	return value >= 65 && value <= 90
}

func casePermutation(sCase []int, sByte []byte) string {
	j := 0
	for i := 0; i < len(sByte); i++ {
		if upperCaseLetter(sByte[i]) {
			if sCase[j] == 0 {
				sByte[i] += 32
			}
			j++
		} else if lowerCaseLetter(sByte[i]) {
			if sCase[j] == 1 {
				sByte[i] -= 32
			}
			j++
		}
	}

	return string(sByte)
}

func pow(n int, p int) int {
	pow := 1
	for ; p > 0; p-- {
		pow *= n
	}

	return pow
}
