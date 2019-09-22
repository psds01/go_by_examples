// arrays : an indexed sequence (of homogenous type?) of a specific length
// type of elements and array's length : both part of array type

package main

import "fmt"

func main() {
	// define a variable `arr` of type `[5]int`
	// length and type of elements are type of that array
	var arr [5]int
	// prints array with 5 `empty int` values i.e. 0
	fmt.Println(arr)
}