package main

import (
	"fmt"
)

func main() {

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

	num := 7
	switch num % 2 {
	case 0:
		fmt.Println(num, "is even")
	default:
		fmt.Println(num, "is odd")
	}

}
