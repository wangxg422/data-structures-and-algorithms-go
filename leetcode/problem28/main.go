package main

import "fmt"

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func strStr(haystack string, needle string) int {
	lenH := len(haystack)
	lenN := len(needle)
	if lenH < lenN {
		return -1
	}

	find := false
	for i := 0; i <= lenH-lenN; i++ {
		find = true
		for j := 0; j < lenN; j++ {
			if haystack[i+j] != needle[j] {
				find = false
				break
			}
		}
		if find {
			return i
		}
	}

	return -1
}

func main() {
	s := "leetcode"
	fmt.Println(strStr(s, "code"))
}
