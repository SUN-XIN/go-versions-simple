package main

import (
	"fmt"
	"sort"
)

type foo struct {
	value int
	name  string
}

func main() {
	ff := []foo{
		foo{
			value: 2,
			name:  "tom",
		},
		foo{
			value: 10,
			name:  "alex",
		},
		foo{
			value: 5,
			name:  "bob",
		},
		foo{
			value: 7,
			name:  "jacky",
		},
	}

	sort.Slice(ff, func(i, j int) bool {
		return ff[i].value < ff[j].value
	})
	fmt.Println("sort.Slice -> By int:", ff)

	sort.Slice(ff, func(i, j int) bool {
		return ff[i].name < ff[j].name
	})
	fmt.Println("sort.Slice -> By string:", ff)
}
