// arrays : an indexed sequence (of homogenous type?) of a specific length
// type of elements and array's length : both part of array type

package main

import "fmt"

func main() {
	// declare a variable `arr` of type `[5]int`
	// length and type of elements are type of that array
	var intArr [4]int
	// prints array with 5 `empty int` values i.e. 0
	fmt.Println("int empty arr:", intArr)

	// declare an empty array of bool
	var boolArr [2]bool
	fmt.Println("bool empty arr:", boolArr)

	// declare an empty arr of floats
	var floatArr [3]float64
	fmt.Println("float empty arr:", floatArr)

	// declare an empty array of elements with type string
	var stringArr [10]string
	fmt.Println(stringArr)

	// set value at an index; index starts at 0 and ends at len(arr)-1
	var arr [10]int
	arr[4] = 100
	fmt.Println("Full array", arr)
	fmt.Println("Value at index 4:", arr[4])
	fmt.Println("length of the array", len(arr))

	// declare and initialize an array of type int and of given len
	evenArr := [5]int{2, 4, 6, 8, 10}
	fmt.Println("array of len 5 of type int: containing only even elements:", evenArr)

	// VAR varName [SHAPE][SHAPE]...[SHAPE]INT
	// shapes should be CONST
	const shapeX = 3
	const shapeY = 5
	var twoD [shapeX][shapeY]int
	for i := 0; i < shapeX; i++ {
		for j := 0; j < shapeY; j++ {
			twoD[i][j] = (j + i) * (i + 1)
		}
	}
	fmt.Println(twoD)

}
