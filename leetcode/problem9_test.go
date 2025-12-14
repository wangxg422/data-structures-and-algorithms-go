package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(-123))
	fmt.Println(isPalindrome(-12321))
	fmt.Println(isPalindrome(12332))
	fmt.Println(isPalindrome(123321))
	fmt.Println(isPalindrome(12321))
}
