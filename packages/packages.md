# Packages
Return to overview of notes: [notes.md](../notes.md)

## Code organization with packages
Code in [main.go](basics/main.go)

- packages are required for every Go file in the project
- package `main` and function `main()` specify the entry point for programs
- "shipping" programs means compiling them to get an executable
- syntax for compiling -- `go build -o name_for_executable location_of_file(s)_to_be_compiled`
  - Example: `go build -o testprogram packages/basics/main.go`
- Combine steps of compiling and running by using this syntax -- `go run location_of_files`
  - Example: `go run packages/basics/main.go`
  - **Note:** This creates a temporary executable, runs the program, then deletes the executable

## Importing and using packages
Code in [main.go](imports/begin/main.go)

- Find packages for use in Go code [here](https://pkg.go.dev)
  - Contains documentation for library and third-party packages
  - Searching `func Println` tells me that any number of trailing arguments can be used (indicated by the three dots before "any"), and the function returns an integer for the number of bytes used as well as any error message produced
  - Includes an example of usage which shows which package needs to be included in order to use the function (`fmt` for the `Println` function)
  - Site allows you to modify and run code right there to see how it changes the outcome.
  - Some library packages may be recognized and automatically added if you try to use them. 