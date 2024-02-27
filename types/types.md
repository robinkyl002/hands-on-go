# Types
Return to overview of notes: [notes.md](../notes.md)

## Variables, types, and scope
Code [here](variables/begin/main.go)

Several ways to declare variables in Go
- `var` keyword followed by the variable name and then the type
  - If you do not assign an initial value, the compiler assigns one
    - Called the zero value, it differs between the types
    - for integer it is the number 0
- `var` keyword, then set it equal to an initial value
  - This does not list the type, but the compiler picks up the type from the initial value
- use variable name, `:=`, then the value
  - Method can only be used inside a function not at the package level
- you can create variables in a function with the same names as variables at the package level
  - These variables can also be different types from those at the package level
  - to access the package level variables, you will have to create a function a the package level that allows the function access

## Constants
Code [here](constants/begin/main.go)

- Constants must be Booleans, strings, numerics, or runes
  - runes (aka characters in every other programming language) are a single character in single quotes
- using `%T` in a string with a variable will output the variable type
  - `%v` outputs the value of a number, `%c` outputs the value of a rune/character
- Constants may not be declared using the short method 

## Type Conversions
Code [here](conversion/begin/main.go)

- different from casting in other languages
  - Creates new memory allocation for converted value
  - Need to use a new variable to hold the converted value

## Pointers
Code [here](pointers/begin/main.go)

- pointers are created using this format
  - `var name *variableTypeToBeUsedWith`
    - i.e. `var a *int` will be a pointer used with integers
- Use `&` to access the memory location of a variable
- Use `*` to access the value attached to a pointer

## Struct basics
Code [here](structs/fields/begin/main.go)

- Struct is a collection of fields
- When creating any custom type such as a struct always start it with the `type` keyword
- after including the name of the struct, put the `struct` keyword
- inside the struct, list any fields needed for the struct
- When creating a struct in a function, set the variable equal to the struct name, then use braces to surround the field declarations
  - Use a colon after the field name
  - put a comma after each field is declared, even the last in the list.

## Struct Methods
Code [here](structs/methods/begin/main.go)

- specify the struct as a receiver 
- When declaring the function, after the `func` keyword, put in parentheses the variable name and then the struct name

## Struct Pointers
Code [here](structs/pointers/begin/main.go)

- To pass a struct by reference into a function declaration, put an asterisk next to the struct name

## Struct Embedding
Code [here](structs/embed/begin/main.go)

- Go doesn't support inheritance
- Embedding is used to compose structs and interfaces
- To embed a parent struct, just add its name to the list of fields under the other struct
- When initializing a struct that embedded another struct in itself, you have to use the structure `parentStruct : parentStruct` then inside the braces initialize the fields for that struct like normal
- overriding is allowed for functions of structs

## Arrays
Code [here](arrays/begin/main.go)

- Declare an array using this format
  - `var name [arraySize]type`
    - i.e. `var a [4]int`
- Arrays cannot be resized
  - If you end up needing a bigger array, you need to create a new array and copy your values from the original to the new.
  - Slices are a more convenient way to work with arrays

## Slices
Code [here](slices/begin/main.go)

- slices dynamically manage the array in the background so you don't have to do the work manually
- The make function works with slices, maps, and channels to allocate a size in an optional capacity
- use `append()` to add items to a slice

## Maps
Initialization Code [here](maps/initialization/begin/main.go)
Lookups Code [here](maps/lookups/begin/main.go)
Mutating Code [here](maps/mutating/begin/main.go)

### Initialization 
- key-value pairs
- Declare in the following format
  - `var mapName map[keyType]valueType`
    - i.e. `var authors map[string]author`
- initialize the map with the `make` function
  - i.e. `authors = make(map[string]author)`
  - Don't need to specify size of the map
- Add a pair to the map using this format
  - `mapName[key] = value`
    - i.e. `authors["tm"] = author{name: "Toni Morrison"}`
- Access a value in the map using its key with this format
  - `mapName[key]`

### Lookups
- If you try to access a value using a key that doesn't exist, no error is thrown
  - It just returns an empty value which can cause bugs
- Use comma okay style to lookup things in maps
  - `a, ok := authors["jr]"]`
  - `ok` is a boolean variable that returns true if the key does exist

### Mutating
- To change the value associated with a key use this format
  - `mapName[key] = newValue`
    - i.e. `authors["tm"] = author{name: "James Baldwin"}`
- Delete key function built in for maps
  - `delete(mapName, key)`