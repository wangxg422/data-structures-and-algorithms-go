package main

import "fmt"

// https://leetcode.com/problems/add-binary/
// 性能较差
func addBinary(a string, b string) string {
	if a == "" && b == "" {
		return ""
	}
	if a == "" && b != "" {
		return b
	}
	if a != "" && b == "" {
		return a
	}

	p1 := len(a) - 1
	p2 := len(b) - 1

	m := 0
	r := make([]rune, 0)
	for p1 >= 0 || p2 >= 0 {
		v1 := '0'
		v2 := '0'
		if p1 >= 0 {
			v1 = rune(a[p1])
		}
		if p2 >= 0 {
			v2 = rune(b[p2])
		}

		if v1 == '0' && v2 == '0' {
			if m == 1 {
				r = append([]rune{'1'}, r...)
				m = 0
			} else {
				r = append([]rune{'0'}, r...)
			}
		} else if v1 == '0' && v2 == '1' {
			if m == 1 {
				r = append([]rune{'0'}, r...)
				m = 1
			} else {
				r = append([]rune{'1'}, r...)
			}
		} else if v1 == '1' && v2 == '0' {
			if m == 1 {
				r = append([]rune{'0'}, r...)
				m = 1
			} else {
				r = append([]rune{'1'}, r...)
			}
		} else if v1 == '1' && v2 == '1' {
			if m == 1 {
				r = append([]rune{'1'}, r...)
				m = 1
			} else {
				r = append([]rune{'0'}, r...)
				m = 1
			}
		}

		p1--
		p2--
	}

	if m == 1 {
		r = append([]rune{'1'}, r...)
	}

	return string(r)
}

func main() {
	fmt.Println(addBinary("11", "11"))
	fmt.Println(addBinary("11", "0"))
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("11", "10"))
	fmt.Println(addBinary("101", "111"))
}
