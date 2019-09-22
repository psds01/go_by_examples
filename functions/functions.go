// simple usage if functions

package main

import "fmt"

// sytnax: FUNC funcName(argName TYPE, ..., )return TYPE {... RETURN XX}

func addTwoNumbers(a int, b int) int {
	// explicit return
	return a + b
}

// broadcast type in case of multiple arguments of the same TYPE
func addThreeNumbers(a, b, c int) int {
	return a + b + c
}

func multiplyTwoNumbers(a, b int) int { return a * b }

func main() {
	// adding two numbers
	fmt.Println("adding 2 and 3 using a function =", addTwoNumbers(2, 3))
	fmt.Println("adding 2, 3 and 5 using a function =", addThreeNumbers(2, 3, 5))
	fmt.Println("Multipling 4 and 5 =", multiplyTwoNumbers(4, 5))

}
