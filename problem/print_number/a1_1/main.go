package main

import (
	"fmt"
	"sync"
)

// 语义有问题，用for range读取channel，但是没有close channel，更推荐使用select

var (
	v  = 0
	mu sync.Mutex
)

var wg sync.WaitGroup

func print(name string, in, out chan struct{}) {
	defer wg.Done()

	for _ = range in {
		if v > 100 {
			out <- struct{}{}
			break
		}

		fmt.Printf("%s: %d\n", name, v)
		mu.Lock()
		v++
		mu.Unlock()
		out <- struct{}{}
	}
}

func main() {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)

	wg.Add(3)

	go print("goroutine1", ch1, ch2)
	go print("goroutine2", ch2, ch3)
	go print("goroutine3", ch3, ch1)

	ch1 <- struct{}{}

	wg.Wait()
}
