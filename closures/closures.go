package main

import "fmt"

// function which returns another function defined inline
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println()
	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())
}
