package main

func binarySearch(s []int, left int, right int, value int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2
	if s[mid] < value {
		return binarySearch(s, mid+1, right, value)
	} else if s[mid] > value {
		return binarySearch(s, left, mid-1, value)
	} else {
		return mid
	}
}

func binarySearch2(s []int, value int) int {
	left := 0
	right := len(s) - 1

	for left <= right {
		mid := left + (right-left)/2
		if s[mid] < value {
			left = mid + 1
		} else if s[mid] > value {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
