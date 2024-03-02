// interfaces/type-assertions/begin/main.go
package main

import "fmt"

func main() {
	// perform a type assertion
	var i interface{} = 1
	// fmt.Println(i.(string))

	var j any = "yellow"
	fmt.Println(j.(string))

	// perform a type assertion and handle the error
	if _, ok := i.(int); !ok {
		fmt.Printf("%v is not an int\n", i)
	} else {
		fmt.Printf("i is %v\n", i)
	}
}
