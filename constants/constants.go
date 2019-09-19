package main

import (
	"fmt"
	"math"
)

// CONST varName TYPE = value
const myConstantString string = "CONSTANT_STRING"

func main() {
	fmt.Println(myConstantString)
	// CONST varName = value
	const bigNumber = 500000000
	fmt.Println("Big number", bigNumber)
	const derivedBigNumber = 3e20 / bigNumber
	fmt.Println("Derived Big Number, scientific", derivedBigNumber)
	fmt.Println("Derived Big Number, int", int64(derivedBigNumber))

	fmt.Println(math.Sin(bigNumber))

}
