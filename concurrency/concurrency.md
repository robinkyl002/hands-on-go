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

- Create a channel using the following syntax
  - `channelName := make(chan type)`
  - Example: `ch := make(chan int)`
  - Leaving off the type creates an unbuffered channel
- Add a parameter to the function you want to use the channel for so that it accepts a channel
- If you know data will only be passed out of the function on the channel, put an arrow pointing to chan
  - i.e. `func sum(nums []int, ch chan<- int)`
- Send data back to the main thread by using this syntax
  - `ch <- data`
- Receive the data in main using this syntax
  - `result := <-ch`
- Use that variable for whatever comes next

## Buffered Channels
Code [here](channel-basics/begin/main.go)

- When creating a channel, if you know you will be passing more than one piece of data, specify the number of data pieces

## Channel Ranging
Code [here](channel-ranging/begin/main.go)

- You can use a `for` loop with a limit of `cap(channel)` to loop through a channel
  - `cap()` can also be used to capture the size of things like arrays and slices
- Random numbers are generated using the `math/rand` package
- Make sure to close the channel when needed to avoid errors

## Channel Select
Code [here](channel-select/begin/main.go)

- When you need to wait on multiple channels at the same time
- `select` allows you to give multiple channels and uses the results from the one that responds first
- The `select` uses `cases`
- 

## Channel Non-blocking
Code [here](channel-non-blocking/begin/main.go)

- Avoiding blocking when sending on or receiving from a channel
- Use `default` case in `select`
- `default` is a catch-all
- Infinite loops are often used with non-blocking channels
  - infinite loop is `for` loop with no parameters

## Sync Package
Code [here](sync/begin/main.go)

- `sync` package provides basic synchronization primitives including
  - Mutual exclusion locks
  - wait group type - helps wait for a collection of goroutines to finish
    - Declare using `sync.WaitGroup`
    - use the `Add()` function to tell the group how many routines are being added
      - You can use add on each pass in the loop or, if you know the number ahead of time, you can just add that number
    - Once a goroutine has finished, call the `Done()` function
      - make sure to use the `defer` keyword in front of the `Done()` call
    - Use the `Wait()` function to make sure that all of the routines run before continuing
- Maps in Go are not thread-safe
  - Potential for data erase because multiple threads may be trying to write to it at the same time
- Dash race flag
  - When running the program in the terminal, place `-race` between `go run` and the file
  - Shows how the goroutines may cause conflicts, should one finish before another
- Use mutual exclusion (`mutex`) to help with these issues
  - Implementing `Lock()` and `Unlock()`
  - Initialize using `sync.Mutex`
  - Place `Lock()` and `Unlock()` after the `Done()` 
  - Make sure to use `defer` with `Unlock()`
