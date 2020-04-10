// go run 13.go
package main

import (
	"fmt"
	"reflect"
)

// This is a function called appendStr that returns another function
// The function it returns is an anonymous function
// with the method signature "func(string) string"
func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := appendStr()
	fmt.Println(reflect.TypeOf(a)) // func(string) string

	// notice how the 't' value is being updated within the function
	// with the new value.
	fmt.Println(a("World"))  // Hello World
	fmt.Println(a("Gopher")) // Hello World Gopher
	fmt.Println(a("!"))      // Hello World Gopher !

	// basically our function has become a stateful function.

}
