package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ronna-s/gceu2026/lessons/l1-goroutines/fileservice"
)

type Client interface {
	GetParts() []*fileservice.Part
	GetPart(p *fileservice.Part) ([]byte, int)
}

func main() {
	fmt.Println(AggergateFile(fileservice.NewClient()))
}

func AggergateFile(client Client) string {
	list := client.GetParts()
	out := make([]string, len(list))

	var wg sync.WaitGroup
	for _, part := range list {
		wg.Go(func() {
			b, idx := client.GetPart(part)
			out[idx] = string(b)
		})
	}
	wg.Wait()

	return strings.Join(out, "")
}
