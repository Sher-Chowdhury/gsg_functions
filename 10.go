// go run 10.go
package main

import (
	"fmt"
	"reflect"
)

// Here, we're creating the 'add' data type.
type add func(a int, b int) int

// This function accepts an input parameter in the form of a function.
// This function must be an anonymous function that requires
// 2 integer as input parameters
// and one integer return parameter
// i.e: "func(a, b int) int"
func simple(varfunction func(a, b int) int) {
	fmt.Println(varfunction(60, 7))
}

func main() {

	// Here we're creating an anonymous function inside var1.
	var var1 add = func(a int, b int) int {
		return a + b
	}
	fmt.Println(reflect.TypeOf(var1)) // main.add

	simple(var1) // 67
	// the input+output parameter signatures need to match, for this to work

	// Note the simple function can only accept the var1 function
	// as it's input parameter, not any input parameters for the var function,
	// e.g. the following will give an error message
	// simple(var1(30, 7))
	// but you can still just do
	fmt.Println(var1(30, 7)) // 37

}

// https://golangbot.com/first-class-functions/
