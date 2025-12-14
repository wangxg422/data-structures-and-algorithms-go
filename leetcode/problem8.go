package main

// https://leetcode.com/problems/string-to-integer-atoi/
 
// To Check
func myAtoi(s string) int {
	var sign rune = ' '
	var result int
	inited := false

	for _, c := range s {
		if c == ' ' {
			if inited {
				break
			}
			continue
		}

		if c == '.' {
			break
		}

		if c >= 'a' && c <= 'z' {
			break
		}
		if c >= 'A' && c <= 'Z' {
			break
		}

		if c == '+' || c == '-' {
			if inited {
				break
			}

			if sign == ' ' {
				sign = c
				continue
			} else {
				break
			}
		}

		if c >= '0' && c <= '9' {
			if c == '0' && !inited {
				result = 0
				inited = true
				continue
			} else {
				d := int(c - '0')
				if !inited {
					result = d
				} else {
					result = result*10 + d
				}
				inited = true
			}
		} else {
			break
		}
	}

	if sign == '-' {
		return -result
	}

	return result
}
