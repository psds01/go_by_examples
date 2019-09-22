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
	fmt.Println("The element at index 4 is", stringSlice[4])

}
