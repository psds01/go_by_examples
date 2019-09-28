// struct:: typed collection of fields
// is this kinda a python class?
// used to create records

package main

import "fmt"

// syntax:
// TYPE structName STRUCT{
// 	field1 TYPE
// 	filed2 TYPE
// 	field3 TYPE
// }

type fish struct {
	water         string
	eatsOtherFish bool
}

type person struct {
	name string
	age  int
}

type animal struct {
	food    string
	numLegs int
	hasTail bool
}

type examResult struct {
	subject string
	marks   int
}

type laptop struct {
	modelName string
	ram       int
	ssdSize   int
}

type cat struct {
	name       string
	color      string
	age        int
	gender     string
	isDomestic bool
}

type car struct {
	name       string
	price      int
	horsePower int
}

func NewCar(name string) *car {
	// NewCar :how to declare a varibale using a struct
	// varName := structName{field:value, filed:value,...}
	c := car{name: name}
	c.price = 150000
	c.horsePower = 10
	return &c
}

func NewPerson(name string) *person {
	p := person{name: name}
	// Answer to everything.
	p.age = 42
	return &p

}

func main() {
	lamborghini := car{name: "lamborghini"}
	fmt.Println(lamborghini)
	lamborghiniNew := car{name: "lamborghini", price: 150000, horsePower: 100}
	fmt.Println(lamborghiniNew)
	ferrari := NewCar("ferrari")
	fmt.Println(ferrari)

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 24})
	fmt.Println(person{age: 24, name: "Alice"})
	fmt.Println(person{name: "Carol"})
	fmt.Println(&person{name: "Dan", age: 25})
	fmt.Println(NewPerson("Emma"))
}
