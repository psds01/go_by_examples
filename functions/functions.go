// simple usage if functions

package main

import "fmt"

// sytnax: FUNC funcName(argName TYPE, ..., )return TYPE {... RETURN XX}

func addTwoNumbers(a int, b int) int {
	// explicit return
	return a + b
}

func subtractTwoNumbers(a, b int) int {
	return a - b
}

// broadcast type in case of multiple arguments of the same TYPE
func addThreeNumbers(a, b, c int) int {
	return a + b + c
}

func multiplyTwoNumbers(a, b int) int {
	return a * b
}

func divideTwoNumbers(a, b int) float64 {
	return float64(a) / float64(b)
}

// return types are int and float!!!!
func multiplyAndDivideTwoNumber(a, b int) (int, float64) {
	return multiplyTwoNumbers(a, b), divideTwoNumbers(a, b)
}

func addAndSubtractTwoNumber(a, b int) (int, int) {
	// Oh yeah!!
	// re-using the functions defined!
	return addTwoNumbers(a, b), subtractTwoNumbers(a, b)
}

// does not return anything, hence no return type is given!
func sumArbitraryNumberOfInts(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

func main() {

	// fmt.Println(2 / 3)
	// fmt.Println(3 / 2)
	// adding two numbers
	fmt.Println("adding 2 and 3 using a function =", addTwoNumbers(2, 3))
	fmt.Println("adding 2, 3 and 5 using a function =", addThreeNumbers(2, 3, 5))
	fmt.Println("Multipling 4 and 5 =", multiplyTwoNumbers(4, 5))

	fmt.Println("Printing addition and difference of 10 and 5")
	// declare and initialize addition and subtraction
	addition, difference := addAndSubtractTwoNumber(10, 5)
	fmt.Println("addition:", addition, "difference:", difference)

	// unlike python, `_` can not be used to get value of the last expression
	_, div := multiplyAndDivideTwoNumber(10, 5)
	fmt.Println("Division of 10 and 5 is", div)

	mul, _ := multiplyAndDivideTwoNumber(10, 5)
	fmt.Println("Multiplication of 10 and 5 is", mul)

	// variadic function
	sumArbitraryNumberOfInts(10)
	sumArbitraryNumberOfInts(10, 20, 40)

	// nums := [5]int{1, 2, 3, 4, 5}
	// WTF: does not accept arr!!
	nums := []int{1, 2, 3, 4, 5}
	sumArbitraryNumberOfInts(nums...)

}
