package main

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
