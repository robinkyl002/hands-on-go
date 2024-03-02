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

## Type Assertions

Code [here](type-assertions/begin/main.go)

## Type Switches

Code [here](type-switch/begin/main.go)