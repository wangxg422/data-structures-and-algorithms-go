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
    for i := 0; i < n/2; i++ {
        minIdx := i
        maxIdx := i
        // 正确的边界是 j < n-i
        for j := i + 1; j < n-i; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
            if arr[j] > arr[maxIdx] {
                maxIdx = j
            }
        }

        // 把最小值换到前面
        if minIdx != i {
            arr[i], arr[minIdx] = arr[minIdx], arr[i]
            // 如果最大值位置就在 i，则刚被换走，需要修正 maxIdx
            if maxIdx == i {
                maxIdx = minIdx
            }
        }

        // 把最大值换到后面
        if maxIdx != n-1-i {
            arr[n-1-i], arr[maxIdx] = arr[maxIdx], arr[n-1-i]
        }
    }
}

func selectSort3(arr []int) {
    left, right := 0, len(arr)-1

    for left < right {
        minIdx, maxIdx := left, left

        for i := left + 1; i <= right; i++ {
            if arr[i] < arr[minIdx] {
                minIdx = i
            }
            if arr[i] > arr[maxIdx] {
                maxIdx = i
            }
        }

        arr[left], arr[minIdx] = arr[minIdx], arr[left]
        // 如果最大值被换走就修正
        if maxIdx == left {
            maxIdx = minIdx
        }

        arr[right], arr[maxIdx] = arr[maxIdx], arr[right]

        left++
        right--
    }
}

