package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// 方法1
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			if i == len(nums)-1 {
				nums = nums[0:i]
			} else {
				nums = append(nums[0:i], nums[i+1:]...)
				i-- // 因为删除了一个元素，i要回退一个，因为之后要自增1
			}
		} else {
			count++
		}
	}

	return count
}

// 方法2
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		} else {
			nums[index] = nums[i]
			index++
		}
	}

	if index > 1 {
		nums = nums[0 : index]
	}

	return index
}

func main() {
	nums := []int{
		1, 1, 1, 2, 2, 3, 3, 3, 3, 4,
	}
	c := removeDuplicates(nums)
	fmt.Println(c)
	fmt.Println(len(nums))
	fmt.Println(nums)
}
