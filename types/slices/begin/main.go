// types/slices/begin/main.go
package main

import "fmt"

func main() {
	// declare a slice string the make function
	names := make([]string, 0)

	// append 3 names to the slice
	names = append(names, "John")
	names = append(names, "Jane")
	names = append(names, "Mary")

	// print the slice
	fmt.Println(names)

	// slice the slice using syntax slice[low:high]
	fmt.Println(names[1:3]) // [Jane Mary] - Everything from the low index to one before the high index
	fmt.Println(names[1:]) // [Jane Mary] - everything from the low index to the end
	fmt.Println(names[:1]) // [John] - everything from the beginning to the index before the high index
	fmt.Println(names[:]) // [John Jane Mary] - Everything in the slice
}
