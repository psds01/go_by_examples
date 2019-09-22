// range "iterates" over different data structures
// range :: an iterator

package main

import "fmt"

func main() {
	// iterate over a slice
	numbers := []int{1, 2, 3}
	// print above numbers one by one
	// need to "declare" range
	// if gives both index and value at the index
	// syntax: FOR index, value := RANGE numbers {...}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// sum the elements in an array
	evenNumbers := [5]int{2, 4, 6, 8, 10}
	sum := 0
	for i, num := range evenNumbers {
		sum += num
		fmt.Println("Sum till", i, "=", sum)
	}

	fmt.Println("Sum of total elements in the array", sum)

}
