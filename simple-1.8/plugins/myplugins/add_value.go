// go build -buildmode=plugin -o add_val.so add_value.go
package main

import (
	"fmt"
	"reflect"
)

func AddInt(x, y int) int {
	fmt.Printf("Add int: %d + %d \n", x, y)
	return x + y
}

func AddString(x, y string) string {
	fmt.Printf("Add string: %s + %s \n", x, y)
	return x + y
}

func Show(x interface{}) {
	fmt.Printf("Show: %v has type %v \n", x, reflect.TypeOf(x))
}
