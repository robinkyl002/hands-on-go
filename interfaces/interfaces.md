# Interfaces
Return to overview of notes: [notes.md](../notes.md)

## Basics - Defining and Implementing

Code [here](basics/begin/main.go)

- Interfaces specify a set of methods that must be implemented
- No need to embed an interface in a struct. Using functions of an interface will tell the program that the struct is pulling from the interface
- A type must implement every method in an interface
- `Stringer` interface in the `fmt` package allows overwriting the standard output format used by Go
  - The only method in `Stringer` is `String()`

## Empty Interface

Code [here](empty/begin/main.go)

- Empty interface is just an interface without any methods
- It allows you to apply functions to variables of (almost?) any type

## Type Assertions

Code [here](type-assertions/begin/main.go)

- Assert a variable as a certain type before working with it
- another way to set a value as an empty interface is to use the `any` keyword
- make the assertion by putting a period next to the variable, then in parentheses the type that the assertion is for
- Be careful with assertions as they can cause problems
  - Use ok to check whether the assertion was done correctly

## Type Switches

Code [here](type-switch/begin/main.go)