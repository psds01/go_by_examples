package main

import "fmt"

func main() {
	if 42%2 == 0 {
		fmt.Println("It's always 42!")
	} else {
		fmt.Println("Except when it's not.")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

}
