package main

import "fmt"

// https://leetcode.com/problems/3sum/
// TO finished
func threeSum(nums []int) [][]int {
	r := [][]int{}
	n := len(nums)
	if n < 3 {
		return r
	}

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					add(r, nums[i], nums[j], nums[k])
				}
			}
		}
	}

	return r
}

func add(r [][]int, v1 int, v2 int, v3 int) {

}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1}))
	fmt.Println(threeSum([]int{-2, -1, 0, 1, 2}))
}
