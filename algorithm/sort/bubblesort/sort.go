package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 1; j < n-i; j++ {
			if arr[j] < arr[j-1] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func bubbleSortFlag(arr []int) {
	n := len(arr)

	isSwapped := false
	for i := 0; i < n-1; i++ {
		isSwapped = false
		for j := 1; j < n-i; j++ {
			if arr[j-1] > arr[j] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
				isSwapped = true
			}
		}
		if !isSwapped {
			fmt.Println("swap is not happened")
			break
		}
	}
}
