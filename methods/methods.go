// package main

// import "fmt"

// // syntax: STRUCT
// type car struct {
// 	numberTires, hp int
// 	name            string
// }

// type rect struct {
// 	width, height float64
// }

// type circle struct {
// 	radius float64
// }

// type square struct {
// 	side float64
// }

// // SYNTAX:
// // FUNC (varName TYPE_OF_VAR) funcName() RETURN TYPE {
// // 	// do something
// // }

// // receiver type of *struct
// func (r *rect) area() float64 {
// 	return r.width * r.height
// }

// func (c circle) area() float64 {
// 	return (c.radius * 22) / 7
// }

// func (s square) area() float64 {
// 	return s.side * s.side
// }

// func main() {
// 	fmt.Println("Hey!")
// 	// declare and initialize a rectangle
// 	r := rect{height: 2, width: 3}
// 	fmt.Println(r.width, r.height)
// 	fmt.Println("Area of rectangle", r.area())

// 	c := circle{3}
// 	fmt.Println(c.radius)
// 	fmt.Println(c.area())

// 	s := square{4}
// 	fmt.Println(s.side)
// 	fmt.Println(s.area())
// }

package main

import (
	"fmt"
)

// make a struct
type employee struct {
	monthlySalary float64
}

// pointer receiver method
// this should be the default way to define methods
// the method has no parameters
// in-place ops possible
// avoid copying and can mutate struct
func (e *employee) ctc() float64 {
	return e.monthlySalary * 12
}

// value receiver method
// NOT recommended
// the method has no parameters
// no in-place support
// copies and mutates struct
func (e employee) taxToPay() float64 {
	return 0.3 * e.monthlySalary * 12
}

func main() {
	// pointer receiver method
	emp := employee{100000}
	fmt.Println(emp.monthlySalary)
	fmt.Println(emp.ctc())
	fmt.Println("\n")

	// value receiver method
	emp1 := employee{100000}
	emp2 := &emp1
	fmt.Println(emp2.monthlySalary)
	fmt.Println(emp2.ctc())
}
