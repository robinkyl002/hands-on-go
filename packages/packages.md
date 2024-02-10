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