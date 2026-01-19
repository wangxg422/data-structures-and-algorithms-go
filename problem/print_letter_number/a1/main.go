package main

import (
	"fmt"
	"sync"
)

func Print() {
	letter := make(chan struct{})
	number := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		i := 1
		for {
			_, ok := <-number
			if !ok {
				return
			}
			fmt.Print(i)
			i++
			letter <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for l := 'A'; l <= 'Z'; l++ {
			<-letter
			fmt.Print(string(l))
			if l == 'Z' {
				close(number)
				return
			}
			number <- struct{}{}
		}
	}()

	number <- struct{}{}
	wg.Wait()
}

func main() {
	Print()
}
