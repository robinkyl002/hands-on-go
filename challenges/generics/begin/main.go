// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring

// non-generic print functions

// type printer interface {
// 	~string | constraints.Integer | ~bool
// }

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

func printAny[T any](a T) {
	fmt.Println(a)
}

// Part 2 sum function refactoring

// sum sums a slice of any type
func sum(numbers []interface{}) interface{} {
	var result float64
	for _, n := range numbers {
		switch n.(type) {
		case int:
			result += float64(n.(int))
		case float32, float64:
			result += n.(float64)
		default:
			continue
		}
	}
	return result
}

type numeric interface {
	constraints.Integer | constraints.Float
}

// func sumAny[T numeric](numbers []interface{}) T {
// 	var result T
// 	for _, n := range numbers {
// 		switch n.(type) {
// 		case int:
// 			result += T(n.(int))
// 		case float32, float64:
// 			result += n.(T)
// 		default:
// 			continue
// 		}
// 	}
// 	return result
// }

// Solution from the instructor
//
func sumAny[T numeric](numbers ...T) T {
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}

func main() {
	// call non-generic print functions
	// printString("Hello")
	// printInt(42)
	// printBool(true)

	// call generic printAny function for each value above
	// printAny("Hello")
	// printAny(42)
	// printAny(true)

	// call sum function
	fmt.Println("result", sum([]interface{}{1, 2, 3}))
	fmt.Println("result", sum([]interface{}{1.2, 2.3, 3.4, 4.5}))

	// call generics sumAny function
	fmt.Println("result:", sumAny(1, 2, 3))
	fmt.Println("result:", sumAny(1.2, 2.3, 3.4, 4.5))
}
