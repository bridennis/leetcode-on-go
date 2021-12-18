package easy_roman_to_integer

/**
https://leetcode.com/problems/roman-to-integer/
*/
func romanToInt(s string) int {
	romanToInt := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	special := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	intNumber := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			chs := string(s[i]) + string(s[i+1])
			if v, ok := special[chs]; ok {
				intNumber += v
				i++
				continue
			}
		}

		if v, ok := romanToInt[string(s[i])]; ok {
			intNumber += v
		}
	}

	return intNumber
}
