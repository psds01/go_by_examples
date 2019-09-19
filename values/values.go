// value types: integer, float, boolean, string.

package main

import "fmt"

func main() {
	// string ops
	fmt.Println("\n\nString operations:--->")

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
	fmt.Println("\n\nInteger operations:--->")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7-3 =", 7-3)
	fmt.Println("1-1 =", 1-1)
	// int * int = int
	fmt.Println("12*4 =", 12*4)
	// int / int = int
	fmt.Println("20/4 = ", 20/4)

	// float ops
	fmt.Println("\n\nFloat operations:--->")
	// int / float = float
	fmt.Println("4/5.0 = ", 4/5.0)
	//float / int = float
	fmt.Println("4.0/5 = ", 4.0/5)
	// float / int = float
	fmt.Println("With `fmt.Println` -> 22.0/7.0 = ", 22.0/7.0)
	// WTF? is fmt a specialized package for printing?
	// float / float = float
	println("With `built-in println` -> 22.0/7.0 = ", 22.0/7.0)
	// fmt.Println("22.0/7.0 =", 22.0/7.0)

	// bool ops
	fmt.Println("\n\nBoolean operations:--->")
	// && = and
	fmt.Println("true && false =", true && false)
	// || = or
	fmt.Println("true || false =", true || false)
	// ! = not
	fmt.Println("!false =", !false)

}
