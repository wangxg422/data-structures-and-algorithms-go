package main

import "fmt"

func do(s []int) {
	// v是切片的副本
	for _, v := range s {
		v += 10
	}
}

func do2(s []int) {
	for i, v := range s {
		s[i] = v + 10
	}
}

func main() {
	var s = []int{0, 1, 2, 3, 4}
	do(s)
	fmt.Println(s)
	do2(s)
	fmt.Println(s)
}
