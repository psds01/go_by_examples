// value types: integer, float, boolean, string.

package main

import "fmt"

func main() {
	// string ops

	// prints "golang"; ends with newline
	fmt.Println("go" + "lang")
	// prints "golang"; no newline at the end
	fmt.Print("go" + "lang")
	// no newline at the end
	fmt.Printf(" is" + " good")

	// Starts printing where the last printX function left
	// NOTE
	fmt.Println("Jaadu hai, nasha hai!")
	fmt.Print("madhahoshiyaan, tujko bhula ab ")
	fmt.Printf("jau kaha\n")

	// int ops
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7-3 =", 7-3)
	fmt.Println("1-1 =", 1-1)
	fmt.Println("12*4 =", 12*4)
	fmt.Println("20/4 = ", 20/4)

}
