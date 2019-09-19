package main

import (
	"fmt"
)

func main() {
	// explicit variable declaration, followed by variable name followed by value assignment
	var a = "variable a"
	fmt.Println("printing variable a :->", a)

	// VAR var1, var2 TYPE = val1, val2
	var int1, int2 int = 1, 2
	fmt.Println("variables define by type broadcasting :->", int1, int2)

	// can I assign value to multiple variables without sepcifying a TYPE?
	var a1, a2 = "a1", "a2"
	fmt.Println("Variables assigned without TYPE broadcasting :->", a1, a2)

	// VAR varName = anyValue
	var boolVar = true
	fmt.Println(boolVar)

	// VAR varName TYPE
	fmt.Println("Any variable defined with type but not value will assume the default value of that type")
	var zeroInt int
	var zeroString string
	var zeroFloat float32
	var zeroBool bool
	fmt.Println("Variables defined with type but not with value :--->\n", "int", zeroInt, "string", zeroString, "flaot", zeroFloat, "bool", zeroBool)
	var emptyString string
	fmt.Println("empty string here:", emptyString, ":ends here")

	// another method of declaration
	// varName := value
	myFruit := "apple"
	fmt.Println(myFruit)

	// varName := value
	myCar := "Ferrari"
	// varName := value
	myBalance := 0
	fmt.Println("Car I want:", myCar, "|| money I have:", myBalance)
	dreamComeTrue := true
	fmt.Println("Do dreams come true?", dreamComeTrue)

}
