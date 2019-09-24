package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func sliceModify(slice []int) {
	for i, elem := range slice {
		if elem%2 == 0 {
			slice[i] = elem * 2
		}
	}
}

func main() {
	i := 1
	fmt.Println("pointer", &i)
	fmt.Println("Original", i)
	zeroval(i)
	fmt.Println("Zeroeval", i)
	zeroptr(&i)
	fmt.Println("Zeroptr", i)
	fmt.Println("pointer", &i)

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice before modifying...", mySlice)
	sliceModify(mySlice)
	fmt.Println("modified slice data...", mySlice)
}
