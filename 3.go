package main

import (
	"fmt"
	"reflect"
)

// Here, all you need to set a return value is specify it's datatype
// which in this case is just string.
func functionWithReturnValue(aName string, aCity string, anAge int) string {

	fmt.Println("'name' is set to:", aName) // 'name' is set to: Tony Stark
	fmt.Println("'city' is set to:", aCity) // 'city' is set to: Manhattan
	fmt.Println("'age' is set to:", anAge)  // 'age' is set to: 3

	return "Avengers member"
}

func main() {

	ironManName := "Tony Stark"
	ironManCity := "Manhattan"
	ironManAge := 3
	result1 := functionWithReturnValue(ironManName, ironManCity, ironManAge)

	fmt.Println(reflect.TypeOf(result1)) // string
	fmt.Println(result1)                 // Avengers member

}

// Also see:
// https://golangbot.com/first-class-functions/
