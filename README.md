# Stopwatch
[![Go Reference](https://pkg.go.dev/badge/github.com/woodywood117/stopwatch.svg)](https://pkg.go.dev/github.com/woodywood117/stopwatch)

Stopwatch is a simple stopwatch library that allows you to start and pause a timer.

## Example
```go
package main

func main() {
    // Create a new stopwatch
    s := stopwatch.New()

    // Start the stopwatch
    s.Start()

    // Do something

    // Pause the stopwatch
    s.Pause()

    // Do something else

    // Resume the stopwatch
    s.Start()

    // Do something

    // Pause the stopwatch again
    s.Pause()

    // Print the elapsed time
    fmt.Println(s.Elapsed())
}
```

You can get the elapsed time in both the running and paused state.