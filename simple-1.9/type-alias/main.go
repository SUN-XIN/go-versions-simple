package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Age  int8
	Name string
}

func main() {
	type Employer = Person

	emp := Employer{
		Age:  10,
		Name: "toto",
	}

	pers := Person{
		Age:  10,
		Name: "toto",
	}

	fmt.Println("Type Person=Employer ? ", emp == pers)

	fmt.Println("TypeOf Person", reflect.TypeOf(pers))
	fmt.Println("TypeOf Employer", reflect.TypeOf(emp))
}
