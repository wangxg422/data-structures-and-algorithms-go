package main

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var arr = make([]int, 10)

	for i := 0; i < 10; i++ {
		arr[i] = 10 - i
	}

	bubbleSort(arr)

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}

func TestBubbleSortFlag(t *testing.T) {
	var arr = make([]int, 10)

	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	bubbleSortFlag(arr)

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}
