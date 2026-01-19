package main

import (
	"fmt"
	"sync"
)

func Print() {
	letter := make(chan bool)
	number := make(chan bool)

	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		defer wait.Done()
		i := 1
		for range number {
			fmt.Print(i)
			i++
			letter <- true
		}
	}()

	go func() {
		defer wait.Done()
		i := 'A'
		for range letter {
			fmt.Print(string(i))
			if i >= 'Z' {
				close(number)
				return
			}
			i++
			number <- true
		}
	}()

	number <- true
	wait.Wait()
}

func main() {
	Print()
}
