package main

import "math"

// https://leetcode.com/problems/sqrtx/
func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}