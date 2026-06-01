package main

import (
	_ "embed"
	"testing"
	"testing/synctest"

	"github.com/ronna-s/gceu2026/lessons/goroutines/fileservice"
)

//go:embed assets/preview.html
var originalFile string

func TestAggergateFile(t *testing.T) {
	_, expectedGoRoutines := fileservice.NewClient().Parts()
	var (
		nGoRoutines int
		output      string
	)

	synctest.Test(t, func(t *testing.T) {
		output = AggergateFile()
		for range expectedGoRoutines {
			if originalFile == output {
				break
			}
			nGoRoutines++
			synctest.Wait()
		}
	})
	if output != originalFile {
		t.Fatalf("Unexpected output from AggergateFile: '%s'", output)
	}
	if expectedGoRoutines != nGoRoutines {
		t.Fatalf("Unexpected number of goroutines invoked. Expected: %d, Saw: %d", expectedGoRoutines, nGoRoutines)
	}

}
