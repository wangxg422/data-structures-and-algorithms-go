package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ProcessA(ctx context.Context) {
	finish := make(chan struct{})
	go func() {
		time.Sleep(time.Second) // 模拟 1 秒的处理逻辑
		finish <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		fmt.Printf("A cancel, err:%s\n", ctx.Err())
	case <-finish:
		fmt.Printf("A finish\n")
	}
}

func ProcessB(ctx context.Context) {
	finish := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second) // 模拟 2 秒的处理逻辑
		finish <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		fmt.Printf("B cancel, err:%s\n", ctx.Err()) // B cancel, err:context deadline exceeded
	case <-finish:
		fmt.Printf("B finish\n")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		ProcessA(ctx)
	}()
	go func() {
		defer wg.Done()
		ProcessB(ctx)
	}()
	wg.Wait()
}
