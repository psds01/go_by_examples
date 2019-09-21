package main

import "fmt"

func main() {
	if 42%2 == 0 {
		fmt.Println("It's always 42!")
	} else {
		fmt.Println("Except when it's not.")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// The scope of this variale is just for the if...else part
	if num := 11; num%2 == 0 {
		fmt.Println("11 is divisible by 2")
	} else if num%5 == 0 {
		fmt.Println("11 is divisible by 5")
	} else {
		fmt.Println("11 is divisible neither by 2 nor by 5")
	}
	// undefined: num
	// fmt.Println(num)

}
