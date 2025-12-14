package main

// https://leetcode.com/problems/reverse-integer/description/

func reverse(x int) int {
	r := 0
	for x != 0 {
		t := x % 10
		r = r*10 + t
		x /= 10
	}
	
	return r
}