// go errors are returned by explicit, separate value.
// TODO: Why this design choice?

package main

import (
	"errors"
	"fmt"
)

// BY CONVENTION, errors are the last return types
// syntax:
// func funcName(arg argType)(RETURN TYPE, ERROR){
// 	return true, errors.New("what")
// }

// a function which return an error on some condition
func f1(arg int) (int, error) {
	if arg == 42 {
		// error string SHOULD not be caps, not should end with punct.
		return -1, errors.New("this can't be the answer to everything")
	}
	// nil is the SINGLETON?
	return arg + 3, nil
}

// what is this type for?
type argError struct {
	arg  int
	prob string
}

// WTF is going on now?
func (e *argError) Error() string {
	return fmt.Sprintf("%d -%s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{42, "this can't be the answer to everything, come on now"}
	}
	return arg + 3, nil
}

func main() {

	// build in way to generate and catch error
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	// custom error type
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	// call f2 with arg = 42
	_, e := f2(42)
	// WTF is this syntax??
	// e.(*argError) ??
	// OK!! creating an instance of error type to use the data from the generated error!!
	// _, myCustomReturnedError := someFunc()
	// arg1, arg2, ... := myCustomReturnedError.(*myCustomErrorType)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
