package ratelimit

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestRateLimiter(t *testing.T) {
	const maxReqs = 10
	interval := time.Second
	// the rate limiter allows up to maximum maxReqs requests/interval
	limiter := NewAtomicRateLimiter(maxReqs)
	_ = limiter

	// test that it doesn't allow more than that, for 2 seconds.
	synctest.Test(t, func(t *testing.T) {
		// Start must be executed inside the bubble
		stop := limiter.Start(interval, maxReqs)
		defer stop()

		for range 5 {
			for range maxReqs {
				if !limiter.Allow() {
					t.Error("expected true got false")
				}
			}
			if limiter.Allow() {
				t.Error("expected false got true")
			}
			time.Sleep(interval)
			synctest.Wait() // ensures that refill happens
		}
	})
}
