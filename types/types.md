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

