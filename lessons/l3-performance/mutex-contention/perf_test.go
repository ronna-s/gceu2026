package mutexcontention_test

import (
	"runtime"
	"sync"
	"testing"
)

func BenchmarkContendedMutex(b *testing.B) {
	var mu sync.Mutex
	var x int

	runtime.SetMutexProfileFraction(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			x++
			mu.Unlock()
		}
	})
}
