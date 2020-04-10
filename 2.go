// go run 2.go
package main

import (
	"fmt"
)

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

func main() {

	ironManName := "Tony Stark"
	ironManCity := "Manhattan"
	ironManAge := 3
	functionWithparameters(ironManName, ironManCity, ironManAge)

}
