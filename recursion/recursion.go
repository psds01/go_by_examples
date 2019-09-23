package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Printing factorial of 7:", fact(7))
	for i := 0; i < 10; i++ {
		fmt.Println("Fibonacci series at index", i, ":", fibonacci(i))
	}
}
