package main

import "fmt"

// https://leetcode.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	space := ' '
	n := len(s)
	max := 0

	for i := n - 1; i >= 0; i-- {
		v := (rune)(s[i])
		if max == 0 && v == space {
			continue
		}

		if max == 0 && v != space {
			max++
			continue
		}

		if max > 0 && v != space {
			max++
		}

		if max > 0 && v == space {
			break
		}
	}

	return max
}

func main() {
	fmt.Println(lengthOfLastWord(""))
	fmt.Println(lengthOfLastWord("   "))
	fmt.Println(lengthOfLastWord("  aa"))
	fmt.Println(lengthOfLastWord("aa"))
	fmt.Println(lengthOfLastWord(" aa "))
	fmt.Println(lengthOfLastWord("aa "))
}
