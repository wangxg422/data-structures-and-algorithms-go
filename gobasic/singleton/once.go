package main

import "sync"

var (
	instance1 *singleton
	once      sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		instance1 = &singleton{}
	})
	return instance1
}
