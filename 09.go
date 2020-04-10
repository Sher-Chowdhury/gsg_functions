// go run 09.go
package main

import (
	"fmt"
	"reflect"
)

// We can create a function in the form a of variable type.
// here we're saying that whenever we create a variable of the type 'add',
// The box for this variable will be designed to store a function
// that takes in 2 integers as input parameters and returns a single integer parameter.
type add func(a int, b int) int

func main() {

	// Here we're creating variable called var1, which is of the 'add' datatype.
	// And we've stored a function inside that variable.
	var var1 add = func(a int, b int) int {
		return a + b
	}
	fmt.Println(reflect.TypeOf(var1)) // main.add

	// now we can use this special syntax to feed input parameters
	// to the funciton that lives inside our var1 function variable.
	s := var1(5, 6)
	fmt.Println(reflect.TypeOf(s)) // int

	fmt.Println("Sum", s) // Sum 11
}

// https://golangbot.com/first-class-functions/

// Since var1 is technically a variable. it means we can write functions that
// accepts var1 as an input parameter. we'll cover that next.
