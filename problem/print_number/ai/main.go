package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		v  = 0
		mu sync.Mutex
	)

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(3)

	print := func(name string, in, out chan struct{}) {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			case <-in:
				mu.Lock()
				if v > 100 {
					mu.Unlock()
					close(done) // 广播退出
					return
				}
				fmt.Printf("%s: %d\n", name, v)
				v++
				mu.Unlock()
				out <- struct{}{}
			}
		}
	}

	go print("goroutine1", ch1, ch2)
	go print("goroutine2", ch2, ch3)
	go print("goroutine3", ch3, ch1)

	ch1 <- struct{}{}
	wg.Wait()
}
