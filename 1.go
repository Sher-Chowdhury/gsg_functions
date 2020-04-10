// go run 1.go
package main

import (
	"fmt"
)

func basicFunction() {
	name := "Peter Parker"
	city := "New York"
	age := 17

	fmt.Println("'name' is set to:", name)
	fmt.Println("'city' is set to:", city)
	fmt.Println("'age' is set to:", age)
}

func main() {

	// Here  we are calling the function into action
	basicFunction()
}

// Also see:
// https://golangbot.com/first-class-functions/
// https://www.callicoder.com/golang-functions/
