package main

import (
	"fmt"
	"testing"
)

func TestBinSearch(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	fmt.Println(binSearch(s, 0, len(s)-1, 7))
	fmt.Println(binSearch(s, 0, len(s)-1, 71))
}

func TestBinSearch2(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	fmt.Println(binSearch2(s, 7))
	fmt.Println(binSearch2(s, 71))
}
