package main

import (
	"fmt"
	"sort"
)

type foo struct {
	value int
	name  string
}

type fooByInt []foo

func (v fooByInt) Len() int           { return len(v) }
func (v fooByInt) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v fooByInt) Less(i, j int) bool { return v[i].value < v[j].value }

type fooByString []foo

func (v fooByString) Len() int           { return len(v) }
func (v fooByString) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v fooByString) Less(i, j int) bool { return v[i].name < v[j].name }

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

	sort.Sort(fooByInt(ff))
	fmt.Println("sort.Sort -> By int:", ff)

	sort.Sort(fooByString(ff))
	fmt.Println("sort.Sort -> By string:", ff)

}
