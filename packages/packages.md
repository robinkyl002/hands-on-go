# Packages
Return to overview of notes: [notes.md](../notes.md)

## Code organization with packages
Code in [basics](basics/main.go)

- packages are required for every Go file in the project
- package `main` and function `main()` specify the entry point for programs
- "shipping" programs means compiling them to get an executable
- syntax for compiling -- `go build -o name_for_executable location_of_file(s)_to_be_compiled`
  - Example: `go build -o testprogram packages/basics/main.go`
- Combine steps of compiling and running by using this syntax -- `go run location_of_files`
  - Example: `go run packages/basics/main.go`
  - **Note:** This creates a temporary executable, runs the program, then deletes the executable

## Importing and using packages
Code in [imports](imports/begin/main.go)

- Find packages for use in Go code [here](https://pkg.go.dev)
  - Contains documentation for library and third-party packages
  - Searching `func Println` tells me that any number of trailing arguments can be used (indicated by the three dots before "any"), and the function returns an integer for the number of bytes used as well as any error message produced
  - Includes an example of usage which shows which package needs to be included in order to use the function (`fmt` for the `Println` function)
  - Site allows you to modify and run code right there to see how it changes the outcome.
  - Some library packages may be recognized and automatically added if you try to use them. 

## Packages and Visibility
Code in [visibility](visibility/main.go)

- Go does not use codewords such as `public` or `private` to restrict access to variables or functions, instead it uses capitalization to achieve that
- You can find the code behind a function by looking it up at https://pkg.go.dev and clicking on the function name. 
- In the code for the `Fprintln` function, it uses a `newPrinter` function which cannot be used outside of its package because it starts with a lowercase letter

## Third-party Packages
Code in [modules](modules/main.go)

- You can tell if there is a module being used by whether there is a go.mod file in the project
  - The file has the module's path up at the top, which is what is used to import it into other packages or programs
  - the `require` code block lists the third-party packages used across the project
  - You can get the path for a package on the Go documentation page by clicking on the "copy path" icon at the top, just above the name of the package
    - Example [here](/packages/Third-party%20package%20for%20Go.png)
  - In terminal -- `go get path` (i.e. `go get github.com/jboursiquot/go-proverbs`)
  - The module tool will retrieve the package if it hasn't already been retrieved and cache it locally
  - When adding or removing modules, use `go mod tidy` which will go through and retrieve or remove packages as needed to keep dependencies clean
