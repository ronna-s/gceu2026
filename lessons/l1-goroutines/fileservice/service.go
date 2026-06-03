package fileservice

import (
	"io/fs"
	"log"
	"slices"
	"time"

	"github.com/ronna-s/gceu2026/lessons/l1-goroutines/assets"
)

type service struct {
	parts []*Part
}

var DefaultService = newService()

func newService() *service {
	var srv service
	entries, err := fs.ReadDir(assets.Parts, "parts")
	if err != nil {
		log.Fatalf("failed to read embedded parts directory: %v", err)
	}

	for _, e := range entries {
		path := "parts/" + e.Name()
		b, err := assets.Parts.ReadFile(path)
		if err != nil {
			log.Fatalf("failed to read file %q: %v", path, err)
		}
		srv.parts = append(srv.parts, &Part{p: append(b, '\n')})
	}
	return &srv
}

// GetParts returns a slice of pointers to parts of a file.
func (s service) GetParts() []*Part {
	time.Sleep(time.Second)
	return s.parts
}

func (s service) GetPart(p *Part) ([]byte, int) {
	time.Sleep(time.Second)
	idx := slices.Index(s.parts, p)
	if idx < 0 {
		panic("unkown part given")
	}
	return []byte(p.p[1:]), int(p.p[0])
}
