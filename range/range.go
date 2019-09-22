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

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	evenSum := 0
	oddSum := 0
	indicesForOddNumber := []int{}
	indicesForEvenNumber := []int{}
	for index, num := range nums {
		if num%2 == 0 {
			evenSum += num
			indicesForEvenNumber = append(indicesForEvenNumber, index)
		} else {
			oddSum += num
			indicesForOddNumber = append(indicesForOddNumber, index)
		}
	}
	fmt.Println("Sum of elements at indices", indicesForEvenNumber, "is", evenSum)
	fmt.Println("Sum of elements at indices", indicesForOddNumber, "is", oddSum)

	// range over a map
	// declare a map
	// myMap := make(map[string]int)
	// declare and initialize a map
	myMap := map[string]string{"a": "a", "b": "b", "c": "c", "d": "d"}
	fmt.Println("range over key, val of", myMap)
	// sytnax: FOR key, val := RANGE myMap {...}
	for key, val := range myMap {
		fmt.Println("key =", key, "value =", val)
	}

	// range over keys
	fmt.Println("range over keys of", myMap)
	// syntax: FOR key := RANGE myMap {...}
	for key := range myMap {
		fmt.Println("Key =", key)
	}

	// range over strings
	// weird:: prints unicode valu
	for index, unicodeValue := range "golang" {
		fmt.Println("Found", unicodeValue, "at index", index)
	}

	// how to print actual characters then??

}
