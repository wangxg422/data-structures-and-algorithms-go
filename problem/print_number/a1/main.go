package main

import (
	"fmt"
	"sync"
)

// 语义有问题，用for range读取channel，但是没有close channel，更推荐使用select

func main() {
	var (
		v  = 0
		mu = sync.Mutex{}
	)

	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	ch3 := make(chan struct{}, 1)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		for range ch1 {
			if v > 100 {
				ch2 <- struct{}{}
				break
			}

			fmt.Printf("goroutine1: %d\n", v)
			mu.Lock()
			v++
			mu.Unlock()
			ch2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for _ = range ch2 {
			if v > 100 {
				ch3 <- struct{}{}
				break
			}

			fmt.Printf("goroutine2: %d\n", v)
			mu.Lock()
			v++
			mu.Unlock()
			ch3 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for _ = range ch3 {
			if v > 100 {
				ch1 <- struct{}{}
				break
			}

			fmt.Printf("goroutine3: %d\n", v)
			mu.Lock()
			v++
			mu.Unlock()
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}

	wg.Wait()
}
