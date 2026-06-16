# Demystifying Go Concurrency

## Clone the repo: github.com/ronna-s/gceu2026

## Exercise - Basic Goroutines:
Implement the body of function `AggergateFile(client Client) string` in `lessons/l1-goroutines/fileaggregator/aggregator.go`. 

`AggergateFile` is given a Client type that makes requests to a file server (which only serves one html file and is notoriously slow), with two methods:

```go
type Client interface {
	GetParts() []*fileservice.Part
	GetPart(p *fileservice.Part) ([]byte, int)
}
```

You task is to aggregate all file parts.

* There’s a test for your convenience that also checks how your application performs. Run it: `go test ./... -run ^TestAggergateFile$`
* When done run `go run lessons/l1-goroutines/fileaggregator/aggregator.go > unknown.html`
* Open unknown.html to see the results.

## Exercise - Implmenting a Basic TCP server

Implementing a basic server in go is very simple.

Implement the function `Serve` which takes two paramaters: `l net.Listener` that listens for new connections and a `handle func(net.Conn) error` to handle each connection.

Requirements:

In a loop call `l.Accept()` which returns a connection (net.Conn) and an error. The loop should stop when Accept() returns an error.

For each connection conn call `handle(conn)` concurrently, if it returns an error, log the error.

`Serve()` must not resume until all accepted connections have been handled.

 ```
 go test ./... -bench=BenchmarkContendedMutex \
  -benchtime=30s \
  -cpuprofile=cpu.out \
  -mutexprofile=mutex.out
```

## Exercise - synctest:


