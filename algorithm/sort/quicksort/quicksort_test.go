package main

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	s := make([]int, 10)
	for i := 0; i < 10; i++ {
		s[i] = 10 - i
	}

	quickSort(s, 0, len(s)-1)
	fmt.Println(s)
}
