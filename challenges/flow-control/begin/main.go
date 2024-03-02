// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("panic recovered", err)
			// fmt.Println("Error:", err)
		}
	}() 

	// use os package to read the file specified as a command line argument
	data, err := os.ReadFile(os.Args[1])

	if err != nil {
		panic("Failed to read file")
		// panic(fmt.Errorf("failed to read file %s", err))
	}

	// fmt.Println("Raw data in bytes")
	// fmt.Println(data)
	// fmt.Println("=====================")
	// fmt.Println()

	// convert the bytes to a string
	myString := string(data[:])

	// fmt.Println("String created from bytes")
	// fmt.Println(myString)
	// fmt.Println("=====================")
	// fmt.Println()

	// initialize a map to store the counts
	m := make(map[string]int)

	// fmt.Println("Map")
	// fmt.Println(m)
	// fmt.Println("=====================")
	// fmt.Println()

	// use the standard library utility package that can help us split the string into words
	wordSlice := strings.Split(myString, " ")
	// using strings.Fields(myString) removes the spaces without having to specify the separator

	// capture the length of the words slice
	m["words"] = len(wordSlice)

	// fmt.Println("Size of words slice: ")
	// fmt.Println(len(wordSlice))
	// fmt.Println(length)
	// fmt.Println("=====================")
	// fmt.Println()

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, w := range wordSlice {
		for _, c := range w {
			if unicode.IsLetter(c) {
				m["letters"] = m["letters"] + 1
				// m["letters"]++
			} else if unicode.IsSymbol(c) || unicode.IsPunct(c) {
				m["symbols"] = m["symbols"] + 1
				// m["symbols"]++
			} else if unicode.IsNumber(c) {
				m["numbers"] = m["numbers"] + 1
				// m["numbers"]++
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(m)
}
