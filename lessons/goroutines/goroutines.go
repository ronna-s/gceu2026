package main

import (
	"io"
	"log"
	"os"
)

func main() {
	const dir = "./assets/parts"
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to read directory '%s'", dir)
	}
	var parts []string
	for _, e := range entries {
		func() {
			f, err := os.Open(e.Name())
			if err != nil {
				log.Fatalf("failed to open file '%s'", e.Name())
			}
			defer f.Close()
			b, err := io.ReadAll(f)
			if err != nil {
				log.Fatalf("failed to read from file '%s'", e.Name())
			}
			parts = append(parts, string(b))
		}()
	}
}

func ConvertImage() {

}
