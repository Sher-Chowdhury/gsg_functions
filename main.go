package main

import (
	"fmt"
)

func main() {

	// Here  we are calling the function into action
	basicFunction()

	ironManName := "Tony Stark"
	ironManCity := "Manhattan"
	ironManAge := 35

	functionWithparameters(ironManName, ironManCity, ironManAge)

}

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
