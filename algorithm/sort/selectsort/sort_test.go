package main

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := make([]int, 10)

	for i, _ := range arr {
		arr[i] = 10 - i
	}

	selectSort(arr)

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}

func TestSelectSort2(t *testing.T) {
	arr := make([]int, 10)

	for i, _ := range arr {
		arr[i] = 10 - i
	}

	selectSort(arr)

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}
