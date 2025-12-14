package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) { 
	fmt.Println(reverse(123))
	fmt.Println(reverse(1234))
	fmt.Println(reverse(54321))
}
