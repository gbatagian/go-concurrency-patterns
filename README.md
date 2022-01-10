# About
This sample project provides some concurrency pattern examples in Go using Goroutines.
There are three design patterns presented:
1. The `waitgroups` pattern.
2. The `multiple channels` pattern.
3. The `worker pool` pattern.

# Run
To execute the sample example cases you need:

1. `Go` installed on your machine
2. Change directory to the project's root directory: `cd .../go-concurrency-patterns`
3. The commands for the different example cases are the following:
    * `go run main.go -mode=waitgroups`
    * `go run main.go -mode=channels`
    * `go run main.go -mode=workerpool`