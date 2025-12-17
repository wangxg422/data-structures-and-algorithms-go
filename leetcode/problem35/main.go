package main

import (
	"fmt"
)

// https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	n := len(nums)

	if target <= nums[0] {
		return 0
	}
	if target >= nums[n-1] {
		return n - 1
	}

	left := 0
	right := n - 1

	for left <= right {
		mid := left + (right-left)/2

		if right-left <= 1 {
			if target == nums[left] {
				return left
			}
			if target == nums[right] {
				return right
			}
			if target > nums[left] && target < nums[right] {
				return left + 1
			}
		}

		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(searchInsert([]int{0, 1, 2, 3, 4, 5, 6}, 4))
	fmt.Println(searchInsert([]int{0, 1, 2, 3, 5, 6}, 4))
	fmt.Println(searchInsert([]int{0, 1, 2, 4, 5, 6}, 0))
	fmt.Println(searchInsert([]int{0, 1, 2, 4, 5, 6}, 7))
}
