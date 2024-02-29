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
If-Else Code [here](if-else/begin/main.go)
Switch Code [here](switch/begin/main.go)

-

## Defer-Panic-Recover
Code [here](defer-panic-recover/begin/main.go)