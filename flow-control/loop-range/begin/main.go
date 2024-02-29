// flow-control/loop-range/begin/main.go
package main

import "fmt"

func main() {
	// initialize a slice of ints
	nums := []int{120,240,360}

	// use for-range to iterate over the slice and print each value
	for i, n := range nums {
		fmt.Println(i, n)
	}

	// declare a map of strings to ints
	myMap := map[string]int{"Hunter": 12, "Jared": 13, "Caleb": 20, "Ryan": 25, "Tom": 50}

	// use for-range to iterate over the map and print each key/value pair
	for k, v := range myMap {
		fmt.Println(k, v)
	}
}
