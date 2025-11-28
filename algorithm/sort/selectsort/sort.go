package main

// 选择排序
func selectSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		if min != i {
			temp := arr[i]
			arr[i] = arr[min]
			arr[min] = temp
		}
	}
}

func selectSort2(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		min := i
		max := n - 1 - i
		for j := i + 1; j <= n-i; j++ {
			if arr[min] > arr[j] {
				min = j
			}

			if arr[max] < arr[j] {
				max = j
			}
		}

		if min != i {
			temp := arr[min]
			arr[min] = arr[i]
			arr[i] = temp
		}

		if max != n-1-i {
			temp := arr[n-1-i]
			arr[n-1-i] = arr[max]
			arr[max] = temp
		}
	}
}
