package main

import (
	_ "embed"
	"slices"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/ronna-s/gceu2026/lessons/l1-goroutines/fileaggregator/fileservice"
)

var delay = time.Second

const file = "012345"

type FakeClient struct {
	t     *testing.T
	parts []*fileservice.Part
}

func NewFakeClient() *FakeClient {
	s := make([]*fileservice.Part, len(file))
	for i := range s {
		s[len(s)-1-i] = &fileservice.Part{}
	}
	return &FakeClient{parts: s}
}

func (c FakeClient) GetParts() []*fileservice.Part {
	return c.parts
}

func (c FakeClient) GetPart(p *fileservice.Part) ([]byte, int) {
	idx := slices.Index(c.parts, p)
	if idx < 0 {
		c.t.Fatalf("part doesn't exist")
		return nil, -1
	}
	time.Sleep(delay)
	return []byte{file[idx]}, idx
}

func TestAggergateFile(t *testing.T) {
	client := NewFakeClient()
	expectedWait := delay
	var (
		done   atomic.Bool
		waited time.Duration
		output string
	)

	synctest.Test(t, func(t *testing.T) {
		go func() {
			output = AggergateFile(client)
			done.Store(true)
		}()

		for !done.Load() {
			time.Sleep(time.Second)
			synctest.Wait()
			waited += time.Second
		}
	})
	if output != file {
		t.Errorf("Unexpected output from AggergateFile: '%s'", output)
	}
	if expectedWait != waited {
		t.Errorf("Waited an unexpected amount of time. Expected: '%v', Waited: '%v'.", expectedWait, waited)
	}
}
