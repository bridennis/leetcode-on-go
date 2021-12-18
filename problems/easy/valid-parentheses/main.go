package easy_valid_parentheses

/**
https://leetcode.com/problems/valid-parentheses/
*/
func isValid(s string) bool {
	stack := make([]string, 0, len(s))

	pars := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, ch := range s {
		str := string(ch)
		parStr := pars[str]

		if len(stack) > 0 && stack[len(stack)-1] == parStr {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, str)
		}
	}

	return len(stack) == 0
}
