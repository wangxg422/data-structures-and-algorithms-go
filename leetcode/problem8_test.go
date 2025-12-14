package main

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	fmt.Println(myAtoi("  123"))
	fmt.Println(myAtoi("  -1234"))
	fmt.Println(myAtoi("  +1234"))
	fmt.Println(myAtoi("  +1234p"))
	fmt.Println(myAtoi("  -042"))
	fmt.Println(myAtoi("  0-1")) // 0
}
