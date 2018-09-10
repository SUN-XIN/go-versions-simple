package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Age  int8
	Name string
}

func (p Person) Explain() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
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

	//////////////////////////////////////////////////////////////////////////////
	////////////////////////////   TypeOf    /////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////
	fmt.Println("TypeOf Person", reflect.TypeOf(pers))
	fmt.Println("TypeOf Employer", reflect.TypeOf(emp))

	//////////////////////////////////////////////////////////////////////////////
	////////////////////////////   Switch    /////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////
	var interfEmployer interface{}
	interfEmployer = emp
	switch t := interfEmployer.(type) {
	// case Person, Employer: -> ERROR! duplicate case Person in type switch
	case Person:
		fmt.Println("switch Employer.(type) -> Person")
	default:
		fmt.Println("switch Employer.(type) -> %v", t)
	}

	var interfPerson interface{}
	interfPerson = pers
	switch t := interfPerson.(type) {
	case Employer:
		fmt.Println("switch Person.(type) -> Employer")
	default:
		fmt.Println("switch Person.(type) -> %v", t)
	}

	//////////////////////////////////////////////////////////////////////////////
	////////////////////////////   Method    /////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////
	fmt.Println("Employer uses Person's method:", emp.Explain())
}
