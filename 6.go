// go run 6.go
package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Here's an "anonymous function" aka "function literal"
	// Anonymous functions are functions without a name.
	// they are the only type of functions that can be nested inside other functions.
	func(message string) {
		fmt.Println(message)
	}("Hi I'm an anonymous function") // Hi I'm an anonymous function

	// notice how we injected arguments with the ending round brackets. That's
	// how it's done since you can't call this function.

	// The above anonymous function only lets you run the anonymous function once due to it's syntax
	// However you can run it multiple times by storing the function in a variable.
	// Here we are effectively creating a variable from an anonymous function
	// we omitted the ending () becuase that can now be injected in later.
	anonyvar := func(message string) {
		fmt.Println(message)
	}

	fmt.Println(reflect.TypeOf(anonyvar)) // func(string)

	// You can then call this variable/function using the normal function syntax
	anonyvar("hello")           // hello
	anonyvar("anonymous world") // anonymous world

}

// Also see:
// https://golangbot.com/first-class-functions/
