# Flow Control
Return to overview of notes: [notes.md](../notes.md)

## Loops

- `for` is used for most loops

### Basic Loop
Code [here](loop-basic/begin/main.go)

- `for` loops do not need parentheses around the conditions
- use `len()` for length/size of a string
- You can use the shorthand to define the variable for the loop

### While Loop
While Loop Code [here](loop-while/begin/main.go)

- Although the function and structure is the same as a while loop in other languages, the `for` keyword is used rather than `while`
- Again, parentheses are not used around the condition

### Range
Code [here](loop-range/begin/main.go)

- Distinguishing feature of Go is the `range` keyword
  - used in conjunction with `for` to iterate over elements of maps, slices, arrays, and channels
- `range` can return both the current index and the value at that index
- In loop, include the variables that will receive the data you are pulling out
  - i.e. `k` and `v` for the key and value from the map

### Infinite
Code [here](loop-infinite/main.go)

- Do not include a condition after `for` then put in the braces what should the loop should do


## Conditionals

### If-Else
Code [here](if-else/begin/main.go)

- if you don't care about the index when iterating through a slice, use a blank identifier, aka an underscore (`_`) for the variable that would hold the index
- `else` must be on the same line as the closing brace of the `if` statement
- In the code `even` is being set as a boolean result of seeing whether dividing the number by 2 gives a remainder of 0
  - If even is true, the code inside the if statement will run. 

### Switch
Code [here](switch/begin/main.go)

- after the `switch` keyword specify what variable will be looked at
- if multiple cases should do the same thing, group them after the `case` keyword, separated by commas
- `default` keyword will apply the code block beneath it to anything that does not match any of the cases described.

## Defer-Panic-Recover
Code [here](defer-panic-recover/begin/main.go)

- `defer` allows a function to run no matter what while another function runs
- Similar to try/catch in other languages
  - Do not use in the same way as try/catch
  - Go doesn't support exception handling
  - Treat errors like values
- You can use as many `defer` statements as you want
  - They run LIFO (last one listed will perform first)
- Anything placed after `defer` statements will also run 
- `panic` is used for exception handling but usually paired with `recover`

## Challenge

- When reading arguments from the command line, the argument count starts at zero (0) after the `go run` 
- To check whether a rune in a string is a number, letter, or symbol, you can use the `unicode` package
- `Split` and `Fields` are part of the `strings` package
- `ReadFile` is part of the `os` package
- You can use `for-range` to iterate over a string
- 