// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "happy"

func main() {
	// create a local variable "val" and assign it an int value
	val := 13

	// print the value of the local variable "val"
	fmt.Printf("%T, local val = %v\n", val, val)

	// print the value of the package-level variable "val"
	printVal()

	// update the package-level variable "val" and print it again
	updateVal("sad")
	printVal()

	// print the pointer address of the local variable "val"
	fmt.Printf("%T, local &val = %v\n", &val, &val)
	// var a *int = &val
	// fmt.Printf("%T, local &val = %v\n", a, a)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	// *a = 100
	*(&val) = 100
	fmt.Printf("local val = %v\n", val)
}

func printVal() {
	fmt.Printf("%T, global val = %v\n", val, val)
}

func updateVal(newString string) {
	val = newString
}
