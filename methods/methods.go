package main

import "fmt"

// syntax: STRUCT
type car struct {
	numberTires, hp int
	name            string
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

// SYNTAX:
// FUNC (varName TYPE_OF_VAR) funcName() RETURN TYPE {
// 	// do something
// }

// receiver type of *struct
func (r *rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return (c.radius * 22) / 7
}

func (s square) area() float64 {
	return s.side * s.side
}

func main() {
	fmt.Println("Hey!")
	// declare and initialize a rectangle
	r := rect{height: 2, width: 3}
	fmt.Println(r.width, r.height)
	fmt.Println("Area of rectangle", r.area())

	c := circle{3}
	fmt.Println(c.radius)
	fmt.Println(c.area())

	s := square{4}
	fmt.Println(s.side)
	fmt.Println(s.area())
}
