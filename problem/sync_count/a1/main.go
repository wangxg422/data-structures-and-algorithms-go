package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 10                      // 最大支持并发
	sum := 100                       // 任务总数
	wg := sync.WaitGroup{}           // 控制主协程等待所有子协程执行完之后再退出。
	ch := make(chan struct{}, count) // 控制任务并发的chan
	// defer close(ch)
	for i := 0; i < sum; i++ {
		wg.Add(1)
		// 申请“并发名额”,当ch满时阻塞，等待释放“名额”
		ch <- struct{}{}
		go func(j int) {
			defer wg.Done()
			// 执行完毕，释放“并发名额”
			defer func() { <-ch }()
			fmt.Println(j)
		}(i)
	}
	wg.Wait()
}
