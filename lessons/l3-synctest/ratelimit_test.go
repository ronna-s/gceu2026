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

	// test that it doesn't allow more than that, for 2 seconds.
	synctest.Test(t, func(t *testing.T) {
		// every time interval add up to maxReqs more requests (refill the quota)
		stop := limiter.Start(interval, maxReqs)
		defer stop()

		synctest.Wait()
		for i := range 10 {
			if !limiter.Allow() {
				t.Errorf("limiter was expected to allow %d requests initially. Attempt: %d", maxReqs, i)
			}
		}
		if limiter.Allow() {
			t.Errorf("limiter was not expected to allow more requests")
		}

		time.Sleep(interval)

		for i := range 10 {
			if !limiter.Allow() {
				t.Errorf("limiter was expected to allow %d requests initially. Attempt: %d", maxReqs, i)
			}
		}
		if limiter.Allow() {
			t.Errorf("limiter was not expected to allow more requests")
		}
	})
}
