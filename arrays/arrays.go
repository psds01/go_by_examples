// arrays : an indexed sequence (of homogenous type?) of a specific length
// type of elements and array's length : both part of array type

package main

import "fmt"

func main() {
	// define a variable `arr` of type `[5]int`
	// length and type of elements are type of that array
	var intArr [4]int
	// prints array with 5 `empty int` values i.e. 0
	fmt.Println("int empty arr:", intArr)

	// empty array of bool
	var boolArr [2]bool
	fmt.Println("bool empty arr:", boolArr)

	// empty arr of floats
	var floatArr [3]float64
	fmt.Println("float empty arr:", floatArr)

	// empty array of elements with type string
	var stringArr [10]string
	fmt.Println(stringArr)

}
