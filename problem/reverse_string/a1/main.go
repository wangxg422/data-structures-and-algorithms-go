package main

import "fmt"

func ReverseString(s string) string {
	str := []rune(s)
	l := len(str)

	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}

	return string(str)
}

func main() {
	fmt.Println(ReverseString("abc123"))
	fmt.Println(ReverseString("abc12"))
}
