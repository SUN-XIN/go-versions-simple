package main

import (
	"fmt"
	"reflect"
)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////   Declaration    //////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
type MyIntType int
type MyIntAlias = int

func (a MyIntType) IntFunc() {
	fmt.Println("IntFunc")
}

/*
// error "cannot define new methods on non-local type int"
func (a MyIntAlias) AliasFunc() {
	fmt.Println("AliasFunc")
}
*/

type Person struct {
	Age  int8
	Name string
}

func (p Person) PersonExplain() string {
	return fmt.Sprintf("Person %s is %d years old", p.Name, p.Age)
}

type Employer = Person

func (e Employer) EmployerExplain() string {
	return fmt.Sprintf("Employer %s is %d years old", e.Name, e.Age)
}

/*
// error -> Person.Explain redeclared in this block
func (e Employer) PersonExplain() string {
	return fmt.Sprintf("Employer/Person %s is %d years old", e.Name, e.Age)
}
*/

func main() {

	var i int = 99
	//var i1 MyIntType = i //error -> MyIntType and int are different types
	var i2 MyIntAlias = i // ok -> MyIntAlias is not a type
	fmt.Println(i2)

	//////////////////////////////////////////////////////////////////////////////
	//////////////////////////     Equality     //////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////

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
	fmt.Println("Employer uses Person's method:", emp.PersonExplain())
	fmt.Println("Employer uses Employer's method:", emp.EmployerExplain())
	fmt.Println("Person uses Employer's method:", pers.EmployerExplain())
}
