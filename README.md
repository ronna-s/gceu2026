# Demystifying Go Concurrency

## Clone the repo: github.com/ronna-s/gceu2026

## Exercise - Aggregate a File:
Implement the body of function `AggergateFile(client Client) string` in `lessons/l1-goroutines/fileaggregator/aggregator.go`. 

`AggergateFile` is given a Client type that makes requests to a file server (which only serves one html file and is notoriously slow), with two methods:

```go
type Client interface {
	GetParts() []*fileservice.Part
	GetPart(p *fileservice.Part) ([]byte, int)
}
```

You task is to aggregate all file parts (the byte slices) concurrently and return a string.

* There’s a test for your convenience that also checks how your application performs. Run it: `go test github.com/ronna-s/gceu2026/lessons/l1-goroutines/fileaggregator -run ^TestAggergateFile$`
* When done run `go run lessons/l1-goroutines/fileaggregator/aggregator.go > unknown.html`
* Open unknown.html to see the results.

## Exercise - Implment a Basic server

Implementing a basic server in go is very simple.

Implement the function `Serve` in `lessons/l1-goroutines/server/server.go` which takes two paramaters: `l net.Listener` that listens for new connections and a `handle func(net.Conn) error` to handle each connection.

Requirements:

In a loop call `l.Accept()` which returns a connection (net.Conn) and an error. The loop should stop when Accept() returns an error.

For each connection conn call `handle(conn)` concurrently, if it returns an error, log the error.

`Serve()` must not resume until all accepted connections have been handled.

* There’s a test for your convenience that also checks how your application performs. Run it: 
```bash
go test --race github.com/ronna-s/gceu2026/lessons/l1-goroutines/server -run ^TestServe$
```

## Exercise - synctest

Write a test for an atomic rate limiter (`AtomicRateLimiter` type) in `lessons/l3-synctest/ratelimit_test.go`

The `AtomicRateLimiter` has one function that needs testing - `Allow() bool`

Test the limiter doesn't allow more than maxReqs per interval (using calls to `Allow() bool`). Run the test for at least two cycles of the time interval.

When ready, run your test.
```bash
go test --race github.com/ronna-s/gceu2026/lessons/l3-synctest -run ^TestRateLimiter$
```

Excercise - Benchmark Parallel Increments

Implement `BenchmarkMixed` in `lessons/l5-performance/contention/contention_test.go`. 

Implement the body of `BenchmarkMixed` in `lessons/l5-performance/contention/contention_test.go` to compare the following operations in contention on an int64 value: 

1. Reading and writing atomically.
2. Reading and writing using a lock (Mutex type).
3. Reading and writing using a read/write mutex (RWMutex type).

To do this the contention package has an `Incr` type with the following methods:

```go
ReadMutex() int64, IncrMutex()
ReadRWMutex() int64, IncrRWMutex()
ReadAtomic() int64, IncrAtomic()
```

For every 10th operation we will perform an increment, other 9 are reads.

To see the results run:
```bash
go test github.com/ronna-s/gceu2026/lessons/l5-performance/contention -run=^$ -bench ^BenchmarkMixed$ -cpu=1,2,4,8,16
```