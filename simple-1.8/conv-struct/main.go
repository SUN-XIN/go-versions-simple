package main

import "fmt"

func main() {
	type T1 struct {
		X int `json:"foo"`
	}
	type T2 struct {
		X int `json:"bar"`
	}

	var v1 T1
	v2 := T2{
		X: 99,
	}

	// go1.6, error -> cannot convert v2 (type T2) to type T1
	// now legal
	v1 = T1(v2)

	fmt.Printf("v1 value: %+v \n", v1)
}
