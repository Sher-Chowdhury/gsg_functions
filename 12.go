// go run 12.go
package main

import (
	"fmt"
	"reflect"
)

// Next we create 3 anonymouse functions, all of which
// having parameter signature to match the
// someNumber's data type signature
var ()

func calculator(action string) func(a, b int) int {
	addVar := func(num1, num2 int) int {
		return num1 + num2
	}

	subtractVar := func(num1, num2 int) int {
		return num1 - num2
	}

	multiplyVar := func(num1, num2 int) int {
		return num1 * num2
	}

	switch action {
	case "addition":
		return addVar
	case "subtraction":
		return subtractVar
	case "multiplication":
		return multiplyVar
	default:
		return addVar
	}
}

func main() {

	calcOption := calculator("multiplication")
	// This calcOption variable can now be used
	// to only perform multiplications.
	fmt.Println(reflect.TypeOf(calcOption)) // func(int, int) int
	fmt.Println(calcOption(3, 5))           // 15

}

// https://golangbot.com/first-class-functions/
