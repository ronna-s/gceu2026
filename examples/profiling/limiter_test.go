package main

import (
	"runtime"
	"testing"
	"time"
)

func BenchmarkRateLimiter(b *testing.B) {
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)

	l := NewRateLimiter(1000, 1000)
	stop := l.StartRefill(1*time.Millisecond, 100)
	defer stop()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = l.Allow()
		}
	})
}
