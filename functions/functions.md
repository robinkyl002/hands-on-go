# Functions
Return to overview of notes: [notes.md](../notes.md)
Notes on packages [here](../packages/packages.md)

Multi-value returns, named returns, naked returns
- Declare functions using the `func` keyword
- Return value is placed after the name of the function
- Go doesn't allow overloading a function. Each function must have a unique name
- variable type goes after the variable name
- variables cannot be declared and given a value in the same line. 
- Function `greetWithNameAndAge` takes in two parameters-name and age- and returns a string called greeting which is defined in the function
  - called a named return
  - naked return is when there is nothing following return, as shown in this function
  - `strconv.Itoa(age)` takes the age which is an integer type and converts it to a string so it can be added to the string
- Multi-value return is when a function can return different things and the last option is an error
  - include the possible return types, including error as the last one
  - if there is no error, and a valid value can be returned, put `nil` as the return for the error