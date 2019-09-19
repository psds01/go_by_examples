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

	for {
		fmt.Println("Crazy for loop, no condition, either break it inside this for loop or return from this loop from the function which has this loop")
		break
	}

	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
