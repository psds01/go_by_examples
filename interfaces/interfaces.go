package main

import (
	"fmt"
	"math"
)

// "Interfaces are named collections of method signatures."
// syntax::
// TYPE interfaceName INTERFACE{
// 	methodName() RETURN TYPE
// 	methodName1() RETURN TYPE
// }

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) bool {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
	return true
}

func main() {
	fmt.Println("ji")

	c := circle{radius: 10}
	r := rect{4, 5}
	measure(c)
	measure(r)

}

// type --> methods --> interface
// obj  --> functional --> OOP
