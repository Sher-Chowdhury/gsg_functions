package main

import (
	"errors"
	"fmt"
	"reflect"
)

func basicFunction() {
	name := "Peter Parker"
	city := "New York"
	age := 17

	fmt.Println("'name' is set to:", name)
	fmt.Println("'city' is set to:", city)
	fmt.Println("'age' is set to:", age)
}

func functionWithparameters(aName string, aCity string, anAge int) {

	fmt.Println("'name' is set to:", aName)
	fmt.Println("'city' is set to:", aCity)
	fmt.Println("'age' is set to:", anAge)
}

/*
Some terminology defintions:
arguements: these are actual values you feed to a function, e.g. 'Tony Stark'
parameters: these are settings a function can accept, e.g. 'aName'

Note if all the parameter types are of the type string, then you can write the
following more shorthand form:

func functionWithparameters(aName, aCity string)
*/

/*
Here, all you need to set a return value is specify it's datatype
which in this case is just string.
*/
func functionWithReturnValue(aName string, aCity string, anAge int) string {

	fmt.Println("'name' is set to:", aName)
	fmt.Println("'city' is set to:", aCity)
	fmt.Println("'age' is set to:", anAge)

	return "Avengers member"
}

/*
This function has 2 return values, the first is a string datatype, and the second is an
'error' datatype.
Also notice that when we have multiple return values we have to use round brackets to
enclose them.
*/
func functionWithMultiReturnValues(aName string, aCity string, anAge int) (string, error) {

	fmt.Println("'name' is set to:", aName)
	fmt.Println("'city' is set to:", aCity)
	fmt.Println("'age' is set to:", anAge)

	return "Super Villain", errors.New("Thanos is not an Avenger")
}

// function with named return values
func functionWithnamedReturnValues(aName string, aCity string, anAge int) (returnVal1 string, returnVal2 error) {

	// The returnvalues are initialised+declared for us, by the function itself
	// since they are part of the function's definition.
	fmt.Println("returnVal1 datatype is:", reflect.TypeOf(returnVal1)) // prints: string //which happens to be empty string
	fmt.Println("returnVal1 datatype is:", reflect.TypeOf(returnVal2)) // error defaults to <nil>

	fmt.Println("'name' is set to:", aName)
	fmt.Println("'city' is set to:", aCity)
	fmt.Println("'age' is set to:", anAge)

	returnVal1 = "Villain"
	returnVal2 = errors.New("Thanos is not an Avenger")

	// Notice we don't have to specify what we are exporting.
	return
}

func main() {

	// Here  we are calling the function into action
	basicFunction()

	ironManName := "Tony Stark"
	ironManCity := "Manhattan"
	ironManAge := 35

	functionWithparameters(ironManName, ironManCity, ironManAge)

	result1 := functionWithReturnValue(ironManName, ironManCity, ironManAge)

	fmt.Println(result1)

	result2, errMessage := functionWithMultiReturnValues("Thanos", "Titan", 150)

	fmt.Println(result2) // Super Villain
	fmt.Println("xxxx")
	fmt.Println(errMessage) // Thanos is not an Avenger
	fmt.Println("xxxx")
	fmt.Println(reflect.TypeOf(errMessage)) // *errors.errorString

	/*
		Here, the underscore is used as a way of saying we don't care about this
		variable so please ignore it. It's a bit like Linux's /dev/null, but for golang
		vairiables. This underscore is referred to as the 'blank identifier'
		https://medium.com/@pcp/blank-identifiers-in-googles-golang-3f6a10e7483a
	*/
	result3, _ := functionWithMultiReturnValues("Ultron", "Earth", 1)
	fmt.Println(result3)

	Val1, Val2 := functionWithnamedReturnValues("Ultron", "Earth", 2)
	fmt.Println(Val1, Val2) // Villain Thanos is not an Avenger

	// Here's an "anonymous function" aka "function literal"
	// Anonymous functions are functions without a name.
	// they are the only typle of functions that can be nested inside other functions.
	func(message string) {
		fmt.Println(message)
	}("Hi I'm an anonymous function") // Hi I'm an anonymous function

	// notice how we injected arguments with the ending round brackets. That's how it's done since you
	// can't call this function.

	// The above anonymous function only lets you run the anonymous function once due to it's syntax
	// here's a way to call an anonymous function lots of times.
	// you are effectively creating a variable from an anonymous function
	anonyvar := func(message string) {
		fmt.Println(message)
	}
	fmt.Println(reflect.TypeOf(anonyvar)) // func(string)
	anonyvar("hello")                     // hello
	anonyvar("anonymous world")           // anonymous world

}
