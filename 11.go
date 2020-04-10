// go run 11.go
package main

import (
	"fmt"
	"reflect"
)

// Here, we're creating the 'add' data type.
type add func(a int, b int) int

// This is the same as before, but now the output parameter is a function based datatype.
func simple() func(a, b int) int {
	// Here we're creating an anonymous function inside a variable called var1.
	var var1 add = func(a int, b int) int {
		return a + b
	}
	return var1
}

func main() {

	var1 := simple()
	fmt.Println(reflect.TypeOf(var1)) // func(int, int) int
	fmt.Println(var1(30, 7))          // 37

}

// https://golangbot.com/first-class-functions/
