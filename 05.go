// go run 5.go
package main

import (
	"errors"
	"fmt"
	"reflect"
)

// function with named return values
// notice we wrote input parameter in slightly shorthand form.
func functionWithnamedReturnValues(aName, aCity string, anAge int) (returnVal1 string, returnVal2 error) {

	// The returnvalues are initialised+declared for us, by the function itself
	// since they are part of the function's definition.
	fmt.Println(reflect.TypeOf(returnVal1)) // string
	fmt.Println(reflect.TypeOf(returnVal2)) // <nil>

	fmt.Println("'name' is set to:", aName)
	fmt.Println("'city' is set to:", aCity)
	fmt.Println("'age' is set to:", anAge)

	// we have to create these variables as part of this function
	// since they have been
	returnVal1 = "Villain"
	returnVal2 = errors.New("Thanos is not an Avenger")

	// Notice we don't have to specify what we are returning.
	return
}

func main() {

	Val1, Val2 := functionWithnamedReturnValues("Ultron", "Earth", 2)
	fmt.Println(Val1, Val2) // Villain Thanos is not an Avenger

}
