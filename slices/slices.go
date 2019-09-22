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
	stringSlice := make([]string, 3)
	fmt.Println(stringSlice)

}
