// go run 4.go
package main

import (
	"fmt"
	"reflect"
)

// variadic functions are function that takes multiple input parameters of the
// same type. That's indicated by the '...' notation

func variadicfunction(values ...int) int {

	fmt.Println(reflect.TypeOf(values)) // []int

	var total int

	// Here we're ignoreing the slicekey.
	for _, slicevalue := range values {
		total = total + slicevalue
	}

	return total
}

func main() {

	totalsum := variadicfunction(10, 20, 5)

	fmt.Println(totalsum) // 35

	// Her's another way.

	numbers := []int{3, 4, 5}

	// this is a specidal notation that feeds a slice (or array) into
	// a variadic function.
	anothersum := variadicfunction(numbers...)

	fmt.Println(anothersum) // 12

}

// variadic functions can accept other input parameters, in which case
// the variadic has to be set last, e.g.
// func variadicfunction(message string, values ...int) int {

// Also see:
// https://yourbasic.org/golang/three-dots-ellipsis/
