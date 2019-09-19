// Only looping construct in GO
package main

import "fmt"

func main() {
	// basic for loop
	fmt.Println("Basic for loop: with single condition...")
	i := 1 // variable definition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// initial;condition;after loop
	fmt.Println("The initial//condition//after for loop...")

	// var can be "defined" in the "initial" part of this loop type
	for j := 7; j > 1; j-- {
		fmt.Println(j)
	}

}
