package main

// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	sL := '('
	sR := ')'
	mL := '['
	mR := ']'
	lL := '{'
	lR := '}'

	left := make([]rune, 0)
	for _, c := range s {
		if c == sL || c == mL || c == lL {
			left = append(left, c)
		} else {
			n := len(left)
			if n == 0 {
				return false
			}
			l := left[n-1]
			left = left[0 : n-1]
			if c == sR && l != sL {
				return false
			}
			if c == mR && l != mL {
				return false
			}
			if c == lR && l != lL {
				return false
			}
		}
	}

	if len(left) != 0 {
		return false
	}

	return true
}
