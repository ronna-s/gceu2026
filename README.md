# Demystifying Go Concurrency

## Clone the repo: github.com/ronna-s/gceu2026

## Exercise 1 - Basic Goroutines:
Implement the body of function `AggergateFile(client Client) string` given a client that makes requests to a file server (which only serves one html file and is notoriously slow), with two methods:

```go
type Client interface {
	GetParts() []*fileservice.Part
	GetPart(p *fileservice.Part) ([]byte, int)
}
```

* There’s a test for your convenience that also checks how your application performs. Run it: `go test ./... -run ^TestAggergateFile$`
* When done run `go run lessons/l1-goroutines/fileaggregator/aggregator.go > unknown.html`
* Open unknown.html to see the results.

## Exercise 2 - Implmenting a Basic TCP server

Implementing a basic TCP server in go is very simple.

 ```
 go test ./... -bench=BenchmarkContendedMutex \
  -benchtime=30s \
  -cpuprofile=cpu.out \
  -mutexprofile=mutex.out
```