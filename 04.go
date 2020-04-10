// go run 4.go
package main

import (
	"errors"
	"fmt"
	"reflect"
)

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

func main() {

	result2, err := functionWithMultiReturnValues("Thanos", "Titan", 150)

	fmt.Println(result2) // Super Villain

	fmt.Println(err) // Thanos is not an Avenger

	fmt.Println(reflect.TypeOf(err)) // *errors.errorString

	/*
		Here, the underscore is used as a way of saying we don't care about this
		variable so please ignore it. It's a bit like Linux's /dev/null, but for golang
		vairiables. This underscore is referred to as the 'blank identifier'
		https://medium.com/@pcp/blank-identifiers-in-googles-golang-3f6a10e7483a
	*/
	result3, _ := functionWithMultiReturnValues("Ultron", "Earth", 1)
	fmt.Println(result3) // Super Villain

}
