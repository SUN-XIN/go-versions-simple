package main

import (
	"fmt"
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("./myplugins/add_val.so")
	if err != nil {
		log.Printf("Failed plugin.Open: %+v", err)
		return
	}

	addInt, err := p.Lookup("AddInt")
	if err != nil {
		log.Printf("Failed Lookup AddInt: %+v", err)
		return
	}

	sumInt := addInt.(func(int, int) int)(11, 2)
	fmt.Println(sumInt)

	addString, err := p.Lookup("AddString")
	if err != nil {
		log.Printf("Failed Lookup AddString: %+v", err)
		return
	}

	sumString := addString.(func(string, string) string)("my", "plugin")
	fmt.Println(sumString)

	showFunc, err := p.Lookup("Show")
	if err != nil {
		log.Printf("Failed Lookup Show: %+v", err)
		return
	}

	showFunc.(func(interface{}))(1.1415926)
}
