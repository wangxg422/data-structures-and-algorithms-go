package main

import "fmt"

// https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 0
	for _, v := range nums {
		if v == val {
			continue
		} else {
			nums[index] = v
			index++
		}
	}

	nums = nums[:index]

	return index
}

func main() {
	nums := []int{1, 2, 3, 3, 4}
	c := removeElement(nums, 3)
	fmt.Println(c)
	fmt.Println(nums)
}
