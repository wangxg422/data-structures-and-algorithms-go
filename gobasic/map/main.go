package main

import "fmt"

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"zhoujielun"}}
	// struct是值传递，通过 map 取出的 struct 是副本，修改副本不会影响原来的 map
	s := m["people"]
	s.name = "wuyanzu"
	fmt.Println(m)
}
