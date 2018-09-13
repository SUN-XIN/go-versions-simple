package main

import (
	"fmt"
	"reflect"

	"github.com/SUN-XIN/go-versions-simple/simple-1.9/type-alias/types"
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

type Student struct {
	Person
}

type User struct {
	Employer
}

type People struct {
	Person
	Employer
}

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

	//////////////////////////////////////////////////////////////////////////////
	////////////////////////     Embedding       /////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////
	var stu Student
	stu.Person.Age = 20
	stu.Person.Name = "bob"
	/*
		// error ->
		// stu.Employer undefined (type Student has no field or method Employer)
		stu.Employer.Age = 20
		stu.Employer.Name = "bob"
	*/
	fmt.Printf("Embedding type Student: %+v -> \n%s\n%s\n\n", stu, stu.EmployerExplain(), stu.PersonExplain())

	var us User
	us.Employer.Age = 30
	us.Employer.Name = "tom"
	/*
		// error ->
		// us.Person undefined (type User has no field or method Person)
		us.Person.Age = 30
		us.Person.Name = "tom"
	*/
	fmt.Printf("Embedding type User: %+v -> \n%s\n%s\n\n", us, us.EmployerExplain(), us.PersonExplain())

	var pp People
	pp.Employer.Age = 11
	pp.Employer.Name = "emp_name"
	pp.Person.Age = 22
	pp.Person.Name = "pers_name"
	fmt.Printf("Embedding type People: %+v \n", pp)
	// error -> ambiguous selector pp.EmployerExplain
	//fmt.Printf("%s \n", pp.EmployerExplain())

	//////////////////////////////////////////////////////////////////////////////
	//////////////////////////     import       //////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////
	// error -> private, can not import
	// var c types.city
	c := types.City{
		Name: "paris",
	}
	fmt.Println("Import type", c)
}
