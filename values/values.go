// value types: integer, float, boolean, string.

package main

import "fmt"

func main() {
	// prints "golang"; the line endswith newline
	fmt.Println("go" + "lang")
	// prints "golang"; no newline at the end
	fmt.Print("go" + "lang")
	// no newline at the end
	fmt.Printf(" is" + " good")

}
