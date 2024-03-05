# Testing
Return to overview of notes: [notes.md](../notes.md)

## Basics
Main file [here](basics/begin/main.go)
Test file [here](basics/begin/main_test.go)

- To do testing, create a file in the same folder as the file to test and name it `[fileName]_test.go`
- Make sure to import the `testing` package
- When creating a test function, start it with "Test" then the name of the function to test
  - i.e. `func TestSum(t *testing.T)`
- When outputting errors, rather than use `fmt` use `t`
- when running a test file through the command line use `go test` rather than `go run`
  - Make sure to use `*.go` to indicate that all go files should be used
- If you want the testing results to be more verbose add `-v` after `go test`
- Use `TestMain` function to run code before and after testing functions
  - make sure to include `m.Run()` in Main or the tests will not run
- If you need setup and teardown to happen inside of a specific function, make sure to do it in that test function, not in TestMain

## Table-driven Tests
Main file [here](table-driven/begin/main.go)
Test file [here](table-driven/begin/main_test.go)

- Testing multiple scenarios/cases at the same time
- Create group of structs with name of the test, input, and expected output

## Benchmarking
Main file [here](benchmarks/begin/main.go)
Test file [here](benchmarks/begin/main_test.go)