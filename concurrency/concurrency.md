# Concurrency
Return to overview of notes: [notes.md](../notes.md)

## Goroutines
Code [here](goroutines/begin/main.go)

- Go supports multi-threading programming with built-in primitives to help make reasoning about programs approachable
- Goroutines are lightweight threads managed by Go runtime
- Thread scheduler is good at scheduling more routines to run across a comparatively smaller number of OS threads
  - Highly efficient use of threads, especially in multiple CPU cores
- Not uncommon to be running thousands of Goroutines on same program
- run a function on its own thread by just putting `go` in front of the function call
  - The main thread pushes the function to a different thread, then goes right back to waiting for the next operation to be done
- Using the `time.Sleep()` function allows the other thread time to run the function and return the result to the main thread
- Channels will allow threads to communicate with each other, helping it to run much smoother than adding a bunch of sleep statements

## Channel Basics
Code [here](channel-basics/begin/main.go)

## Buffered Channels
Code [here]()

## Channel Ranging
Code [here](channel-ranging/begin/main.go)

## Channel Select
Code [here](channel-select/begin/main.go)

## Channel Non-blocking
Code [here](channel-non-blocking/begin/main.go)

## Sync Package
Code [here](sync/begin/main.go)