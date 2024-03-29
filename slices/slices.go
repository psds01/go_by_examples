// a powerful interface to sequences

// TYPEd by the elements they contain, not by length.
// Q: Then how do we get the length of a slice?

package main

import "fmt"

func main() {
	// use builtin make to create a slice

	// declare a slice::
	// mySlice := make([]TYPE, LENGTH)
	// LENGTH is the len of the slice but it's not TYPEd by len
	// The type of the slice is `[]INT`
	stringSlice := make([]string, 7)
	fmt.Println(stringSlice)

	fmt.Println("Setting elements to slice:", stringSlice)
	// set values::
	// syntax : slice[index] = value
	stringSlice[0] = "f"
	stringSlice[1] = "e"
	stringSlice[2] = "r"
	stringSlice[3] = "r"
	stringSlice[4] = "a"
	stringSlice[5] = "r"
	stringSlice[6] = "i"
	fmt.Println("All done. Here's the slice:", stringSlice)

	// get elements by index
	// syntax: slice[index]
	fmt.Println("The element at index 4 is:", stringSlice[4])

	fmt.Println("Current length of the slice:", len(stringSlice))

	// builtin : append
	// syntax : var = append(var, element)
	// variable length append: var = append(var, element1, element2, ...)
	stringSlice = append(stringSlice, ":")
	fmt.Println(stringSlice)
	stringSlice = append(stringSlice, "-", ">")
	fmt.Println(stringSlice)
	stringSlice = append(stringSlice, "I", " ", "w", "a", "n", "t")
	fmt.Println(stringSlice)

	// COPY:
	// syntax : COPY(toCopy, fromCopy)
	copiedSlice := make([]string, len(stringSlice))
	copy(copiedSlice, stringSlice)
	fmt.Println("Print copied slice:", copiedSlice)

	// what if I don't give it the exact length of fromCopy?
	copiedSliceNoLen := make([]string, 2)
	copy(copiedSliceNoLen, stringSlice)
	fmt.Println("It will copy only till the length to toCopy:", copiedSliceNoLen)

	// slice operator
	fmt.Println(stringSlice)
	slc := stringSlice[2:5]
	fmt.Println("2 to 5:", slc)
	fmt.Println("till 5:", stringSlice[:5])
	fmt.Println("from 3:", stringSlice[3:])

	// just like array definition except there is no length
	// declare and initialize a slice
	// just like except the LENGTH TYPE info
	// arr: arrName := [SIZE]TYPE{element,...}
	// slc : slcName := []TYPE{element,...}
	mySlice := []int{1, 1, 2, 3, 5, 8, 13, 21}
	fmt.Println(mySlice)

	// can we declare slice with VAR?
	var varSlice []int
	fmt.Println(varSlice)
	varSlice = append(varSlice, 1)
	varSlice = append(varSlice, 2)
	varSlice = append(varSlice, 3)
	fmt.Println(varSlice)
	fmt.Println(varSlice[:2])
	fmt.Println("Declared slice like a var!! slice notation supported as well!")

	// what will happen if I use slice notation on an arr?
	var myArr [5]int
	fmt.Println(myArr)
	fmt.Println("Trying to slice an arr: ", myArr[1:4])
	fmt.Println("Array also supports slice notation")

	// multi-dimensional slice
	const outerSize = 3
	const stride = 2
	twoDSlice := make([][]int, outerSize)
	for i := 0; i < outerSize; i++ {
		innerSize := i + stride
		// assign a new empty slice to ith element of twoD
		twoDSlice[i] = make([]int, innerSize)
		for j := 0; j < innerSize; j++ {
			// assign value
			twoDSlice[i][j] = i*2 + j*3
		}
	}
	fmt.Println(twoDSlice)
}
