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

}
