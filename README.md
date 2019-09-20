```bash
go get github.com/Sher-Chowdhury/gsg_structs
go run github.com/Sher-Chowdhury/gsg_structs
```

also: 

```bash
$ export PATH=$PATH:~/go/bin/
$ gsg_structs

```



## Nested functions

There is no such thing as nested fuctions, the closest thing to it is [function literals](https://golang.org/ref/spec#Function_literals). 


Normally you shouldn't need to nest functions, but you can nest function using a different syntax, for example lets say you have:


```go
package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {

	fmt.Println(add(1, 1))

}


// the above outputs:
// 2
```

No you can nest the 'add' function inside the the main function like this:


```go
package main

import (
	"fmt"
)

func main() {

    add := func(x, y int) int {
	   		return x + y
	  }
    
	  fmt.Println(add(1, 1))

}

// the above outputs:
// 2
```

An example of where this is used is [recursively get all filepaths from a directory tree](https://www.google.com/search?q=golange+recursively+get+all+files+in+a+directory&oq=golange+recursively+get+all+files+in+a+directory&aqs=chrome..69i57j0.23531j0j7&sourceid=chrome&ie=UTF-8)
