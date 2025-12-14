package main

// https://leetcode.com/problems/palindrome-number/description/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	v := 0
	t := x
	for t != 0 {
		v = v*10 + t%10
		t /= 10
	}

	return x == v
}
