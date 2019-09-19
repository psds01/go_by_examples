package main

import (
	"fmt"
	"math"
)

// CONST can be used wherever VAR can be.

// CONST varName TYPE = value
const myConstantString string = "CONSTANT_STRING"

func main() {
	fmt.Println(myConstantString)
	// CONST varName = value
	const bigNumber = 500000000
	fmt.Println("Big number", bigNumber)
	// CONST expressions perform arithematics operations with arbitrary proportion.
	const derivedBigNumber = 3e20 / bigNumber
	fmt.Println("Derived Big Number, scientific, arbitrary precision", derivedBigNumber)
	fmt.Println("Derived Big Number, int, type given", int64(derivedBigNumber))

	// TYPE assignment by context: bcoz math.Sin expects FLOAT, bigNumber get TYPE float
	fmt.Println(math.Sin(bigNumber))

	// VAR === CONST

}
