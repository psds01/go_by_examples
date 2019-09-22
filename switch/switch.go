package main

import (
	"fmt"
)

func main() {

	// simple switch
	car := "ferrari"
	fmt.Print("I want ")
	switch car {
	case "lamborghini":
		fmt.Println("lamborghini")
	case "honda":
		fmt.Println("honda")
	case "ferrari":
		fmt.Println("ferrari!!!")
	}

	// switch with default
	num := 7
	switch num % 2 {
	case 0:
		fmt.Println(num, "is even")
	default:
		fmt.Println(num, "is odd")
	}

	// switch without an expression
	// same as if
	whatDoWeWant := "car"
	switch {
	case whatDoWeWant == "bike":
		fmt.Println("No, bike is cheap!!")
	case whatDoWeWant == "private jet":
		fmt.Println("Well, I won't mind!!")
	case whatDoWeWant == "car":
		fmt.Println("Oh, yeah... a car")
	}

	// type switch
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println(i, ": I am a bool.")
		case int:
			fmt.Println(i, ": I am an int.")
		default:
			fmt.Println(i, ": Don't know type", t)

		}
	}
	//
	whatAmI(2)
	whatAmI(true)
	whatAmI("Hello")

}
