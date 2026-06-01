package main

import (
	"fmt"
	"strings"

	"github.com/ronna-s/gceu2026/lessons/goroutines/fileservice"
)

func main() {
	fmt.Println(AggergateFile())
}
func AggergateFile() string {
	client := fileservice.NewClient()
	iter, size := client.Parts()

	parts := make([]string, size)

	for part := range iter {
		idx := int(part[0])
		parts[idx] = string(part[1:])
	}

	return strings.Join(parts, "\n")
}
