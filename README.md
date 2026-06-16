# Demystifying Go Concurrency

Welcome to the GopherCon Europe workshop on Go concurrency.

In this workshop, you’ll work through a series of hands-on exercises covering:
- goroutines
- concurrent servers
- deterministic testing with `testing/synctest`
- contention and benchmarking

## Prerequisites

Before starting, make sure you have:

- Go `1.26` or newer
- Git
- A terminal and editor/IDE of your choice

## Setup

Clone the repository:

```bash
git clone https://github.com/ronna-s/gceu2026
cd gceu2026
```

Verify your environment:

```bash
go version
go test ./...
```


## Exercise: Aggregate a File Concurrently

### What you'll do
Implement `AggergateFile(client Client) string` in `lessons/l1-goroutines/fileaggregator/aggregator.go`.

The file server is intentionally slow and returns file parts in random order. Your goal is to fetch the parts concurrently, reassemble them in the correct order, and return the final file contents as a string.

### Success criteria
- all parts are fetched concurrently
- parts are reassembled in the correct order
- the provided test passes

### Run the test

```bash
go test ./lessons/l1-goroutines/fileaggregator -run ^TestAggergateFile$
```

### Optional: inspect the result

```bash
go run ./lessons/l1-goroutines/fileaggregator/aggregator.go > unknown.html
```

Open `unknown.html` in your browser to inspect the reconstructed file.


## Exercise: Implement a Basic Concurrent Server

Implement `Serve` in `lessons/l1-goroutines/server/server.go`.

### Requirements
1. Repeatedly call `l.Accept()`.
2. Stop accepting new connections when `Accept()` returns an error.
3. Handle each accepted connection concurrently by calling `handle(conn)`.
4. If `handle(conn)` returns an error, log it.
5. Do not return from `Serve()` until all accepted connections have finished processing.

### Run the test

```bash
go test --race ./lessons/l1-goroutines/server -run ^TestServe$
```

## Exercise 3: Test a Rate Limiter with `synctest`

In this exercise, you’ll use `testing/synctest` to write a deterministic test for time-based behavior.

Write a test for `AtomicRateLimiter` in `lessons/l3-synctest/ratelimit_test.go`.

### Goal
Verify that `Allow() bool` does not permit more than `maxReqs` requests per interval, and that the quota refills correctly across at least two intervals.

### Run the test

```bash
go test --race ./lessons/l3-synctest -run ^TestRateLimiter$
```


## Exercise: Benchmark Contention Strategies

In this exercise, you’ll compare three approaches to concurrent access to a shared `int64` value:

1. atomic operations
2. `sync.Mutex`
3. `sync.RWMutex`

Implement `BenchmarkMixed` in `lessons/l5-performance/contention/contention_test.go`.

### Goal
Measure mixed read/write contention where:
- every 10th operation is a write
- the other 9 operations are reads

### Run the benchmark

```bash
go test ./lessons/l5-performance/contention -run=^$ -bench ^BenchmarkMixed$ -cpu=1,2,4,8,16
```

## Need help?

If you get stuck:
- re-read the exercise requirements carefully
- run the provided tests frequently
- use the benchmark/test output as feedback
- ask a workshop facilitator for a hint before looking for a full solution