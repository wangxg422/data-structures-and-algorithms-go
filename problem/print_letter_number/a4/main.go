package main

import "fmt"

func Print() {
	number := make(chan bool)
	letter := make(chan bool)
	done := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	go func() {
		j := 'A'
		for {
			select {
			case <-letter:
				fmt.Print(string(j))
				if j >= 'Z' {
					done <- true
					return
				}

				j++
				number <- true
			}
		}
	}()

	number <- true

	for {
		select {
		case <-done:
			return
		}
	}
}

func main() {
	Print()
}
